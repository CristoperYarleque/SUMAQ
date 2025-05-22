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
)

type InfoEntrepreneursDriverInterface interface {
	CreateInfoEntrepreneur(w http.ResponseWriter, r *http.Request)
	GetInfoEntrepreneurs(w http.ResponseWriter, r *http.Request)
}
type infoEntrepreneursDriver struct {
	InfoEntrepreneursService services.InfoEntrepreneursServiceInterface
	BaseDriverInterface
}

func NewInfoEntrepreneursDriver(db database.Connections) InfoEntrepreneursDriverInterface {
	return &infoEntrepreneursDriver{
		InfoEntrepreneursService: services.NewInfoEntrepreneursService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func NewInfoEntrepreneursDriverWithServices(se services.InfoEntrepreneursServiceInterface) InfoEntrepreneursDriverInterface {
	return &infoEntrepreneursDriver{
		InfoEntrepreneursService: se,
		BaseDriverInterface:      &baseDriver{},
	}
}

func (c *infoEntrepreneursDriver) CreateInfoEntrepreneur(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var bodyInfoEntrepreneur services.BodyInfoEntrepreneur
	if err := json.NewDecoder(r.Body).Decode(&bodyInfoEntrepreneur); err != nil {
		c.handleError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
		return
	}

	// Validaciones básicas
	if bodyInfoEntrepreneur.Type == "" || bodyInfoEntrepreneur.Url == "" || bodyInfoEntrepreneur.EntrepreneurId == 0 {
		c.handleError(w, errors.New("all fields (type, url, entrepreneurId) are required"), http.StatusBadRequest)
		return
	}

	err := c.InfoEntrepreneursService.CreateInfoEntrepreneur(ctx, bodyInfoEntrepreneur)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusCreated)
}

func (c *infoEntrepreneursDriver) GetInfoEntrepreneurs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	entrepreneurIdStr := r.URL.Query().Get("entrepreneurId")
	entrepreneurType := r.URL.Query().Get("type")
	if entrepreneurIdStr == "" || entrepreneurType == "" {
		c.handleError(w, errors.New("entrepreneurId and type are required"), http.StatusBadRequest)
		return
	}

	entrepreneurId, err := strconv.Atoi(entrepreneurIdStr)
	if err != nil {
		c.handleError(w, errors.New("entrepreneurId must be a valid integer"), http.StatusBadRequest)
		return
	}

	filterInfoEntrepreneurs := services.FilterInfoEntrepreneurs{
		EntrepreneurId: entrepreneurId,
		Type:           entrepreneurType,
	}

	infoEntrepreneurs, err := c.InfoEntrepreneursService.GetInfoEntrepreneurs(ctx, filterInfoEntrepreneurs)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, infoEntrepreneurs, http.StatusOK)
}
