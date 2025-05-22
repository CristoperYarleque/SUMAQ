package models

import (
	"context"
	"fmt"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type BodyInfoEntrepreneur struct {
	EntrepreneurId int
	Type           string
	Url            string
	Description    string
}

type InfoEntrepreneurs struct {
	Id          int
	Url         string
	Description string
}

type FilterInfoEntrepreneurs struct {
	EntrepreneurId int
	Type           string
}

type InfoEntrepreneursModelInterface interface {
	CreateInfoEntrepreneur(ctx context.Context, bodyInfoEntrepreneur BodyInfoEntrepreneur, querier interfaces.SQLQuerier) error
	GetInfoEntrepreneurs(ctx context.Context, filterInfoEntrepreneurs FilterInfoEntrepreneurs, querier interfaces.SQLQuerier) ([]*InfoEntrepreneurs, error)
}
type infoEntrepreneursModel struct {
	Db *interfaces.SQLConnInterface
}

func NewInfoEntrepreneursModel() InfoEntrepreneursModelInterface {
	return &infoEntrepreneursModel{}
}

func (c *infoEntrepreneursModel) CreateInfoEntrepreneur(ctx context.Context, bodyInfoEntrepreneur BodyInfoEntrepreneur, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		bodyInfoEntrepreneur.EntrepreneurId,
		bodyInfoEntrepreneur.Type,
		bodyInfoEntrepreneur.Url,
		bodyInfoEntrepreneur.Description,
	}

	query := `
		INSERT INTO info_entrepreneur (entrepreneur_id, type, url, description)
		VALUES (?, ?, ?, ?)
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

	fmt.Printf("CreateInfoEntrepreneur - Affected lines %d - RequestID %v \n", rowsAffected, ctx.Value("requestId"))
	return nil
}

func (c *infoEntrepreneursModel) GetInfoEntrepreneurs(ctx context.Context, filterInfoEntrepreneurs FilterInfoEntrepreneurs, querier interfaces.SQLQuerier) ([]*InfoEntrepreneurs, error) {
	var (
		limit string
		where string
	)

	queryParams := []interface{}{
		filterInfoEntrepreneurs.Type,
	}

	if filterInfoEntrepreneurs.Type == "info" {
		where = " AND entrepreneur_id = ? "
		limit = " LIMIT 1 "
		queryParams = append(queryParams, filterInfoEntrepreneurs.EntrepreneurId)
	} else {
		where = ""
		limit = " LIMIT 3 "
	}

	query := `
		SELECT 
		id, 
		url, 
		description 
		FROM info_entrepreneur
		WHERE type = ?
		` + where + `
		ORDER BY id DESC
		` + limit + `
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

	infoEntrepreneurs := make([]*InfoEntrepreneurs, 0)
	for rows.Next() {
		infoEntrepreneur := InfoEntrepreneurs{}
		err := rows.Scan(
			&infoEntrepreneur.Id,
			&infoEntrepreneur.Url,
			&infoEntrepreneur.Description,
		)
		if err != nil {
			log.Println("Error Scan:", err.Error())
			return nil, err
		}
		infoEntrepreneurs = append(infoEntrepreneurs, &infoEntrepreneur)
	}

	return infoEntrepreneurs, nil
}
