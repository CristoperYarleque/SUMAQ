package v1

import (
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

func (config ConnectionConfig) LoginRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewLoginDriver(config.SumaqDB)

	r.Post("/", driver.Login)

	return r
}
