package services

import (
	"context"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
)

type Categories struct {
	CategoryId int
	Name       string
}

type CategoriesServiceInterface interface {
	GetCategories(ctx context.Context) ([]*Categories, error)
}

type categoriesService struct {
	categoriesModel models.CategoriesModelInterface
	getConn         func(ctx context.Context) (*interfaces.SQLInterface, error)
}

func NewCategoriesService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) CategoriesServiceInterface {
	return &categoriesService{
		categoriesModel: models.NewCategoriesModel(),
		getConn:         fn,
	}
}

type CategoriesServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.CategoriesModelInterface
}

func NewCategoriesServiceWithModel(arg CategoriesServiceWithModelArguments) CategoriesServiceInterface {
	return &categoriesService{
		categoriesModel: arg.Vw,
		getConn:         arg.Fn,
	}
}

func (c *categoriesService) GetCategories(ctx context.Context) ([]*Categories, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	categories, err := c.categoriesModel.GetCategories(ctx, db)
	if err != nil {
		log.Printf("error Model GetCategories: %v. Request ID: %v\n", err, ctx.Value("requestId"))
		return nil, err
	}

	categoriesResponse := make([]*Categories, len(categories))

	for i, category := range categories {
		categoriesResponse[i] = &Categories{
			CategoryId: category.CategoryId,
			Name:       category.Name,
		}
	}

	return categoriesResponse, nil

}
