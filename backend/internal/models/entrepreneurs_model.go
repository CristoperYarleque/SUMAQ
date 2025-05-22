package models

import (
	"context"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type Entrepreneurs struct {
	Id    int
	Name  string
	Email string
}

type FilterEntrepreneurs struct {
	CategoryId int
}

type EntrepreneursModelInterface interface {
	GetEntrepreneurs(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs, querier interfaces.SQLQuerier) ([]*Entrepreneurs, error)
}
type entrepreneursModel struct {
	Db *interfaces.SQLConnInterface
}

func NewEntrepreneursModel() EntrepreneursModelInterface {
	return &entrepreneursModel{}
}

func (c *entrepreneursModel) GetEntrepreneurs(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs, querier interfaces.SQLQuerier) ([]*Entrepreneurs, error) {
	queryParams := []interface{}{
		filterEntrepreneurs.CategoryId,
	}

	query := `
		SELECT 
		u.id, 
		u.name, 
		u.email 
		FROM users u
		INNER JOIN products p ON p.entrepreneur_id = u.id
		WHERE p.category_id = ?
		AND u.role = "entrepreneur"
		GROUP BY u.id
	`

	stmt, err := querier.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[FATAL]", "Error prepare Query:", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, queryParams...)
	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	entrepreneurs := make([]*Entrepreneurs, 0)
	for rows.Next() {
		entrepreneur := Entrepreneurs{}
		err := rows.Scan(
			&entrepreneur.Id,
			&entrepreneur.Name,
			&entrepreneur.Email,
		)
		if err != nil {
			log.Println("Error Scan:", err.Error())
			return nil, err
		}
		entrepreneurs = append(entrepreneurs, &entrepreneur)
	}

	return entrepreneurs, nil
}
