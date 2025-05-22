package v1

import (
	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

type ConnectionConfig struct {
	SumaqDB database.Connections
}

func (config ConnectionConfig) CategoriesRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewCategoriesDriver(config.SumaqDB)

	r.Get("/", driver.GetCategories)

	return r
}
