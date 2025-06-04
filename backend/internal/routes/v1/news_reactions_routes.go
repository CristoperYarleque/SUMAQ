package v1

import (
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

func (config ConnectionConfig) NewsReactionsRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewNewsReactionsDriver(config.SumaqDB)

	r.Post("/", driver.UpsertReaction)
	r.Get("/", driver.GetReactionInfo)
	r.Delete("/", driver.DeleteReaction)

	return r
}
