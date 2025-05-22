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
	GetInfoEntrepreneurs(ctx context.Context, filterInfoEntrepreneurs FilterInfoEntrepreneurs, querier interfaces.SQLQuerier) (*InfoEntrepreneurs, error)
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

func (c *infoEntrepreneursModel) GetInfoEntrepreneurs(ctx context.Context, filterInfoEntrepreneurs FilterInfoEntrepreneurs, querier interfaces.SQLQuerier) (*InfoEntrepreneurs, error) {
	queryParams := []interface{}{
		filterInfoEntrepreneurs.EntrepreneurId,
		filterInfoEntrepreneurs.Type,
	}

	query := `
		SELECT 
		id, 
		url, 
		description 
		FROM info_entrepreneur
		WHERE entrepreneur_id = ?
		AND type = ?
		ORDER BY id DESC
		LIMIT 1
	`

	rows := querier.QueryRowContext(ctx, query, queryParams...)

	infoEntrepreneurs := &InfoEntrepreneurs{}
	err := rows.Scan(
		&infoEntrepreneurs.Id,
		&infoEntrepreneurs.Url,
		&infoEntrepreneurs.Description,
	)

	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return nil, err
	}

	return infoEntrepreneurs, nil
}
