package drivers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/services"
)

type EntrepreneursDriverInterface interface {
	GetEntrepreneurs(w http.ResponseWriter, r *http.Request)
}
type entrepreneursDriver struct {
	EntrepreneursService services.EntrepreneursServiceInterface
	BaseDriverInterface
}

func NewEntrepreneursDriver(db database.Connections) EntrepreneursDriverInterface {
	return &entrepreneursDriver{
		EntrepreneursService: services.NewEntrepreneursService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func NewEntrepreneursDriverWithServices(se services.EntrepreneursServiceInterface) EntrepreneursDriverInterface {
	return &entrepreneursDriver{
		EntrepreneursService: se,
		BaseDriverInterface:  &baseDriver{},
	}
}

func (c *entrepreneursDriver) GetEntrepreneurs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categoryIdStr := r.URL.Query().Get("categoryId")
	if categoryIdStr == "" {
		c.handleError(w, errors.New("categoryId is required"), http.StatusBadRequest)
		return
	}

	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.handleError(w, errors.New("categoryId must be a valid integer"), http.StatusBadRequest)
		return
	}

	filterEntrepreneurs := services.FilterEntrepreneurs{
		CategoryId: categoryId,
	}

	entrepreneurs, err := c.EntrepreneursService.GetEntrepreneurs(ctx, filterEntrepreneurs)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, entrepreneurs, http.StatusOK)
}
