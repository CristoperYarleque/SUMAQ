package drivers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/services"
)

type NewsDriverInterface interface {
	CreateNews(w http.ResponseWriter, r *http.Request)
	GetNews(w http.ResponseWriter, r *http.Request)
}
type newsDriver struct {
	NewsService services.NewsServiceInterface
	BaseDriverInterface
}

func NewNewsDriver(db database.Connections) NewsDriverInterface {
	return &newsDriver{
		NewsService: services.NewNewsService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func NewNewsDriverWithServices(se services.NewsServiceInterface) NewsDriverInterface {
	return &newsDriver{
		NewsService:         se,
		BaseDriverInterface: &baseDriver{},
	}
}

func (c *newsDriver) CreateNews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var bodyNews services.BodyNews
	if err := json.NewDecoder(r.Body).Decode(&bodyNews); err != nil {
		c.handleError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
		return
	}

	// Validaciones básicas
	if bodyNews.Title == "" || bodyNews.Content == "" || bodyNews.PublishedAt == "" {
		c.handleError(w, errors.New("all fields (title, content, published_at) are required"), http.StatusBadRequest)
		return
	}

	err := c.NewsService.CreateNews(ctx, bodyNews)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusCreated)
}

func (c *newsDriver) GetNews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")
	if startDate == "" || endDate == "" {
		c.handleError(w, errors.New("startDate and endDate are required"), http.StatusBadRequest)
		return
	}

	filterNews := services.FilterNews{
		StartDate: startDate,
		EndDate:   endDate,
	}

	news, err := c.NewsService.GetNews(ctx, filterNews)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, news, http.StatusOK)
}
