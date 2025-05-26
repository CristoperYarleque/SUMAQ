package models

import (
	"context"
	"fmt"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type Promotion struct {
	PromotionId int
	Description string
	Url         string
	StartDate   string
	EndDate     string
}

type BodyPromotion struct {
	Description string
	StartDate   string
	EndDate     string
}

type FilterPromotions struct {
	StartDate string
	EndDate   string
}

type PromotionsModelInterface interface {
	CreatePromotion(ctx context.Context, bodyPromotion BodyPromotion, querier interfaces.SQLQuerier) error
	GetPromotions(ctx context.Context, filterPromotions FilterPromotions, querier interfaces.SQLQuerier) ([]*Promotion, error)
}
type promotionsModel struct {
	Db *interfaces.SQLConnInterface
}

func NewPromotionsModel() PromotionsModelInterface {
	return &promotionsModel{}
}

func (c *promotionsModel) CreatePromotion(ctx context.Context, bodyPromotion BodyPromotion, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		bodyPromotion.Description,
		bodyPromotion.StartDate,
		bodyPromotion.EndDate,
	}

	query := `
		INSERT INTO promotions (description, start_date, end_date)
		VALUES (?, ?, ?)
	`

	stmt, err := querier.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[FATAL]", "Error prepare Query:", err.Error())
		return err
	}
	defer stmt.Close()

	r, err := stmt.ExecContext(ctx, queryParams...)
	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		log.Printf("Error getting affected rows: %v", err)
		return err
	}

	fmt.Printf("CreatePromotion - Affected lines %d - RequestID %v \n", rowsAffected, ctx.Value("requestId"))
	return nil
}

func (c *promotionsModel) GetPromotions(ctx context.Context, filterPromotions FilterPromotions, querier interfaces.SQLQuerier) ([]*Promotion, error) {
	queryParams := []interface{}{
		filterPromotions.StartDate,
		filterPromotions.EndDate,
	}

	query := `
		SELECT 
		id, 
		description, 
		start_date,
		end_date,
		url
		FROM promotions
		WHERE start_date >= ? 
		AND end_date <= ? 
		ORDER BY 1 DESC;	
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

	promotions := make([]*Promotion, 0)
	for rows.Next() {
		promotion := Promotion{}
		err := rows.Scan(
			&promotion.PromotionId,
			&promotion.Description,
			&promotion.StartDate,
			&promotion.EndDate,
			&promotion.Url,
		)
		if err != nil {
			log.Println("Error Scan:", err.Error())
			return nil, err
		}
		promotions = append(promotions, &promotion)
	}

	return promotions, nil
}
