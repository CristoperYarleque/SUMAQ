package models

import (
	"context"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type Categories struct {
	CategoryId int
	Name       string
}

type CategoriesModelInterface interface {
	GetCategories(ctx context.Context, querier interfaces.SQLQuerier) ([]*Categories, error)
}
type categoriesModel struct {
	Db *interfaces.SQLConnInterface
}

func NewCategoriesModel() CategoriesModelInterface {
	return &categoriesModel{}
}

func (c *categoriesModel) GetCategories(ctx context.Context, querier interfaces.SQLQuerier) ([]*Categories, error) {
	query := `
			SELECT 
			id, 
			name 
			FROM categories 
			ORDER BY 1 DESC;
			`
	stmt, err := querier.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[FATAL]", "Error prepare Query:", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	categories := make([]*Categories, 0)
	for rows.Next() {
		category := Categories{}
		err := rows.Scan(
			&category.CategoryId,
			&category.Name,
		)
		if err != nil {
			log.Println("Error Scan:", err.Error())
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}
