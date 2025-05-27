package drivers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/services"
	"github.com/go-chi/chi/v5"
)

type EntrepreneursDriverInterface interface {
	GetEntrepreneurs(w http.ResponseWriter, r *http.Request)
	UpdateUrlEntrepreneur(w http.ResponseWriter, r *http.Request)
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

func (c *entrepreneursDriver) UpdateUrlEntrepreneur(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	entrepreneurIdStr := chi.URLParam(r, "id")
	if entrepreneurIdStr == "" {
		c.handleError(w, errors.New("id is required"), http.StatusBadRequest)
		return
	}

	entrepreneurId, err := strconv.Atoi(entrepreneurIdStr)
	if err != nil {
		c.handleError(w, errors.New("id must be a valid integer"), http.StatusBadRequest)
		return
	}

	var bodyEntrepreneurUpdate services.BodyEntrepreneurUpdate
	if err := json.NewDecoder(r.Body).Decode(&bodyEntrepreneurUpdate); err != nil {
		c.handleError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
		return
	}

	err = c.EntrepreneursService.UpdateUrlEntrepreneur(ctx, entrepreneurId, bodyEntrepreneurUpdate)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusOK)

}
