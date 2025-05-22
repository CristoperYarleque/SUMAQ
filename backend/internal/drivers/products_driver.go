package drivers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/services"
	"github.com/go-chi/chi/v5"
)

type ProductsDriverInterface interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
}
type productsDriver struct {
	ProductsService services.ProductsServiceInterface
	BaseDriverInterface
}

func NewProductsDriver(db database.Connections) ProductsDriverInterface {
	return &productsDriver{
		ProductsService: services.NewProductsService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func NewProductsDriverWithServices(se services.ProductsServiceInterface) ProductsDriverInterface {
	return &productsDriver{
		ProductsService:     se,
		BaseDriverInterface: &baseDriver{},
	}
}

func (c *productsDriver) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var bodyProduct services.BodyProduct
	if err := json.NewDecoder(r.Body).Decode(&bodyProduct); err != nil {
		c.handleError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
		return
	}

	// Validaciones básicas
	if bodyProduct.Name == "" || bodyProduct.Price == 0 || bodyProduct.CategoryId == 0 || bodyProduct.EntrepreneurId == 0 {
		c.handleError(w, errors.New("all fields (name, price, category_id, entrepreneur_id) are required"), http.StatusBadRequest)
		return
	}

	err := c.ProductsService.CreateProduct(ctx, bodyProduct)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusCreated)
}

func (c *productsDriver) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categoryIdStr := r.URL.Query().Get("categoryId")
	entrepreneurIdStr := r.URL.Query().Get("entrepreneurId")

	if entrepreneurIdStr == "" {
		c.handleError(w, errors.New("entrepreneurId is required"), http.StatusBadRequest)
		return
	}
	entrepreneurId, _ := strconv.Atoi(entrepreneurIdStr)

	categoryId := 0
	if categoryIdStr != "" {
		categoryId, _ = strconv.Atoi(categoryIdStr)
	}

	filterProducts := services.FilterProducts{
		CategoryId:     categoryId,
		EntrepreneurId: entrepreneurId,
	}

	products, err := c.ProductsService.GetProducts(ctx, filterProducts)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, products, http.StatusOK)
}

func (c *productsDriver) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	if id == "" {
		c.handleError(w, errors.New("id is required"), http.StatusBadRequest)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.handleError(w, errors.New("id must be a valid integer"), http.StatusBadRequest)
		return
	}

	var bodyProduct services.BodyProductUpdate
	if err := json.NewDecoder(r.Body).Decode(&bodyProduct); err != nil {
		c.handleError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
		return
	}

	if bodyProduct.Name == "" || bodyProduct.Price == 0 || bodyProduct.CategoryId == 0 {
		c.handleError(w, errors.New("all fields (name, price, category_id) are required"), http.StatusBadRequest)
		return
	}

	err = c.ProductsService.UpdateProduct(ctx, productId, bodyProduct)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusOK)
}

func (c *productsDriver) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	if id == "" {
		c.handleError(w, errors.New("id is required"), http.StatusBadRequest)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.handleError(w, errors.New("id must be a valid integer"), http.StatusBadRequest)
		return
	}

	err = c.ProductsService.DeleteProduct(ctx, productId)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusOK)
}

func (c *productsDriver) GetProductById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	if id == "" {
		c.handleError(w, errors.New("id is required"), http.StatusBadRequest)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.handleError(w, errors.New("id must be a valid integer"), http.StatusBadRequest)
		return
	}

	product, err := c.ProductsService.GetProductById(ctx, productId)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, product, http.StatusOK)
}
