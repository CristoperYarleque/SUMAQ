package services

import (
	"context"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
)

type BodyProduct struct {
	Name           string
	Description    string
	Price          float64
	Url            string
	CategoryId     int
	EntrepreneurId int
}

type BodyProductUpdate struct {
	Name        string
	Price       float64
	Url         string
	CategoryId  int
	Description string
}

type Products struct {
	ProductId    int
	Name         string
	Price        float64
	Url          string
	Description  string
	CategoryId   int
	CategoryName string
}

type FilterProducts struct {
	CategoryId     int
	EntrepreneurId int
}

type ProductsServiceInterface interface {
	CreateProduct(ctx context.Context, bodyProduct BodyProduct) error
	GetProducts(ctx context.Context, filterProducts FilterProducts) ([]*Products, error)
	UpdateProduct(ctx context.Context, productId int, bodyProductUpdate BodyProductUpdate) error
	DeleteProduct(ctx context.Context, productId int) error
	GetProductById(ctx context.Context, productId int) (*Products, error)
}

type productsService struct {
	productsModel models.ProductsModelInterface
	getConn       func(ctx context.Context) (*interfaces.SQLInterface, error)
}

func NewProductsService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) ProductsServiceInterface {
	return &productsService{
		productsModel: models.NewProductsModel(),
		getConn:       fn,
	}
}

type ProductsServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.ProductsModelInterface
}

func NewProductsServiceWithModel(arg ProductsServiceWithModelArguments) ProductsServiceInterface {
	return &productsService{
		productsModel: arg.Vw,
		getConn:       arg.Fn,
	}
}

func (c *productsService) CreateProduct(ctx context.Context, bodyProduct BodyProduct) error {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	bodyProductModel := models.BodyProduct{
		Name:           bodyProduct.Name,
		Description:    bodyProduct.Description,
		Price:          bodyProduct.Price,
		Url:            bodyProduct.Url,
		CategoryId:     bodyProduct.CategoryId,
		EntrepreneurId: bodyProduct.EntrepreneurId,
	}

	err = c.productsModel.CreateProduct(ctx, bodyProductModel, db)
	if err != nil {
		return err
	}

	return nil
}

func (c *productsService) GetProducts(ctx context.Context, filterProducts FilterProducts) ([]*Products, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	filterProductsModel := models.FilterProducts{
		CategoryId:     filterProducts.CategoryId,
		EntrepreneurId: filterProducts.EntrepreneurId,
	}

	products, err := c.productsModel.GetProducts(ctx, filterProductsModel, db)
	if err != nil {
		log.Printf("error Model GetProducts: %v. Request ID: %v\n", err, ctx.Value("requestId"))
		return nil, err
	}

	productsResponse := make([]*Products, len(products))

	for i, product := range products {
		productsResponse[i] = &Products{
			ProductId:    product.ProductId,
			Name:         product.Name,
			Price:        product.Price,
			Url:          product.Url,
			Description:  product.Description,
			CategoryId:   product.CategoryId,
			CategoryName: product.CategoryName,
		}
	}

	return productsResponse, nil

}

func (c *productsService) UpdateProduct(ctx context.Context, productId int, bodyProductUpdate BodyProductUpdate) error {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	bodyProductUpdateModel := models.BodyProductUpdate{
		Name:        bodyProductUpdate.Name,
		Price:       bodyProductUpdate.Price,
		Url:         bodyProductUpdate.Url,
		CategoryId:  bodyProductUpdate.CategoryId,
		Description: bodyProductUpdate.Description,
	}

	err = c.productsModel.UpdateProduct(ctx, productId, bodyProductUpdateModel, db)
	if err != nil {
		return err
	}

	return nil
}

func (c *productsService) DeleteProduct(ctx context.Context, productId int) error {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	err = c.productsModel.DeleteProduct(ctx, productId, db)
	if err != nil {
		return err
	}

	return nil
}

func (c *productsService) GetProductById(ctx context.Context, productId int) (*Products, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	product, err := c.productsModel.GetProductById(ctx, productId, db)
	if err != nil {
		return nil, err
	}

	productResponse := &Products{
		ProductId:   product.ProductId,
		Name:        product.Name,
		Price:       product.Price,
		Url:         product.Url,
		Description: product.Description,
		CategoryId:  product.CategoryId,
	}

	return productResponse, nil
}
