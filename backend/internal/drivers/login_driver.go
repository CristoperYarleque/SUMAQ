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

type LoginDriverInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
}
type loginDriver struct {
	LoginService services.LoginServiceInterface
	BaseDriverInterface
}

func NewLoginDriver(db database.Connections) *loginDriver {
	return &loginDriver{
		LoginService: services.NewLoginService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func (c *loginDriver) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var bodyLogin services.BodyLogin
	if err := json.NewDecoder(r.Body).Decode(&bodyLogin); err != nil {
		c.handleError(w, errors.New("invalid JSON"), http.StatusBadRequest)
		return
	}

	if bodyLogin.Email == "" || bodyLogin.Password == "" || bodyLogin.Role == "" {
		c.handleError(w, errors.New("email, password and role are required"), http.StatusBadRequest)
		return
	}

	login, err := c.LoginService.LoginUser(ctx, bodyLogin)
	if err != nil {
		c.handleError(w, err, http.StatusUnauthorized)
		return
	}

	c.handleSuccess(w, login, http.StatusOK)
}
