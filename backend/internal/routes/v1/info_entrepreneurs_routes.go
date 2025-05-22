package v1

import (
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

func (config ConnectionConfig) InfoEntrepreneursRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewInfoEntrepreneursDriver(config.SumaqDB)

	r.Get("/", driver.GetInfoEntrepreneurs)
	r.Post("/", driver.CreateInfoEntrepreneur)

	return r
}
