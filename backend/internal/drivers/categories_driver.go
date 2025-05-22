package drivers

import (
	"context"
	"net/http"

	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/services"
)

type CategoriesDriverInterface interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
}
type categoriesDriver struct {
	CategoriesService services.CategoriesServiceInterface
	BaseDriverInterface
}

func NewCategoriesDriver(db database.Connections) CategoriesDriverInterface {

	return &categoriesDriver{
		CategoriesService: services.NewCategoriesService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func NewCategoriesDriverWithServices(se services.CategoriesServiceInterface) CategoriesDriverInterface {
	return &categoriesDriver{
		CategoriesService:   se,
		BaseDriverInterface: &baseDriver{},
	}
}

func (c *categoriesDriver) GetCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categories, err := c.CategoriesService.GetCategories(ctx)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, categories, http.StatusOK)
}
