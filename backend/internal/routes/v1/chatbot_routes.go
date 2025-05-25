package v1

import (
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

func (config ConnectionConfig) ChatbotRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewChatbotDriver(config.SumaqDB)

	r.Get("/", driver.GetChatbot)
	r.Post("/", driver.CreateChatbot)

	return r
}
