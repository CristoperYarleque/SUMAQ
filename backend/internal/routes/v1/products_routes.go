package v1

import (
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

func (config ConnectionConfig) ProductsRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewProductsDriver(config.SumaqDB)

	r.Get("/", driver.GetProducts)
	r.Post("/", driver.CreateProduct)
	r.Get("/{id}", driver.GetProductById)
	r.Put("/{id}", driver.UpdateProduct)
	r.Delete("/{id}", driver.DeleteProduct)

	return r
}
