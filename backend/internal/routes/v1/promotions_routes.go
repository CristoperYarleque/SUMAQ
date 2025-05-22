package v1

import (
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

func (config ConnectionConfig) PromotionsRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewPromotionsDriver(config.SumaqDB)

	r.Get("/", driver.GetPromotions)
	r.Post("/", driver.CreatePromotion)

	return r
}
