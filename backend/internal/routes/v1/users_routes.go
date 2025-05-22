package v1

import (
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

func (config ConnectionConfig) UsersRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewUsersDriver(config.SumaqDB)

	r.Post("/", driver.CreateUser)

	return r
}
