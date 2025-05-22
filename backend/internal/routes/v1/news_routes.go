package v1

import (
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

func (config ConnectionConfig) NewsRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewNewsDriver(config.SumaqDB)

	r.Get("/", driver.GetNews)
	r.Post("/", driver.CreateNews)

	return r
}
