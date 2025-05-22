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

type PromotionsDriverInterface interface {
	CreatePromotion(w http.ResponseWriter, r *http.Request)
	GetPromotions(w http.ResponseWriter, r *http.Request)
}
type promotionsDriver struct {
	PromotionsService services.PromotionsServiceInterface
	BaseDriverInterface
}

func NewPromotionsDriver(db database.Connections) PromotionsDriverInterface {
	return &promotionsDriver{
		PromotionsService: services.NewPromotionsService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func NewPromotionsDriverWithServices(se services.PromotionsServiceInterface) PromotionsDriverInterface {
	return &promotionsDriver{
		PromotionsService:   se,
		BaseDriverInterface: &baseDriver{},
	}
}

func (c *promotionsDriver) CreatePromotion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var bodyPromotion services.BodyPromotion
	if err := json.NewDecoder(r.Body).Decode(&bodyPromotion); err != nil {
		c.handleError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
		return
	}

	// Validaciones básicas
	if bodyPromotion.Description == "" || bodyPromotion.StartDate == "" || bodyPromotion.EndDate == "" {
		c.handleError(w, errors.New("all fields (description, start_date, end_date) are required"), http.StatusBadRequest)
		return
	}

	err := c.PromotionsService.CreatePromotion(ctx, bodyPromotion)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusCreated)
}

func (c *promotionsDriver) GetPromotions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")
	if startDate == "" || endDate == "" {
		c.handleError(w, errors.New("startDate and endDate are required"), http.StatusBadRequest)
		return
	}

	filterPromotions := services.FilterPromotions{
		StartDate: startDate,
		EndDate:   endDate,
	}

	promotions, err := c.PromotionsService.GetPromotions(ctx, filterPromotions)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, promotions, http.StatusOK)
}
