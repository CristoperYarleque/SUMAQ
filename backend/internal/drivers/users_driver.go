package drivers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/services"
)

type UsersDriverInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
type usersDriver struct {
	UsersService services.UsersServiceInterface
	BaseDriverInterface
}

func NewUsersDriver(db database.Connections) UsersDriverInterface {
	return &usersDriver{
		UsersService: services.NewUsersService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func NewUsersDriverWithServices(se services.UsersServiceInterface) UsersDriverInterface {
	return &usersDriver{
		UsersService:        se,
		BaseDriverInterface: &baseDriver{},
	}
}

func (c *usersDriver) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var bodyUser services.BodyUser
	if err := json.NewDecoder(r.Body).Decode(&bodyUser); err != nil {
		c.handleError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
		return
	}

	// Validaciones básicas
	if bodyUser.Name == "" || bodyUser.Email == "" || bodyUser.Password == "" || bodyUser.Role == "" {
		c.handleError(w, errors.New("all fields (name, email, password, role) are required"), http.StatusBadRequest)
		return
	}

	err := c.UsersService.CreateUser(ctx, bodyUser)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusCreated)
}
