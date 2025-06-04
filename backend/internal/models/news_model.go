package models

import (
	"context"
	"fmt"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type News struct {
	NewsId      int
	Title       string
	Content     string
	Url         string
	PublishedAt string
}

type BodyNews struct {
	Title       string
	Content     string
	PublishedAt string
}

type FilterNews struct {
	StartDate string
	EndDate   string
}

type NewsModelInterface interface {
	CreateNews(ctx context.Context, bodyNews BodyNews, querier interfaces.SQLQuerier) error
	GetNews(ctx context.Context, filterNews FilterNews, querier interfaces.SQLQuerier) ([]*News, error)
}

type newsModel struct {
	Db *interfaces.SQLConnInterface
}

func NewNewsModel() NewsModelInterface {
	return &newsModel{}
}

func (c *newsModel) CreateNews(ctx context.Context, bodyNews BodyNews, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		bodyNews.Title,
		bodyNews.Content,
		bodyNews.PublishedAt,
	}

	query := `
		INSERT INTO news (title, content, published_at)
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

	fmt.Printf("CreateNews - Affected lines %d - RequestID %v \n", rowsAffected, ctx.Value("requestId"))
	return nil
}

func (c *newsModel) GetNews(ctx context.Context, filterNews FilterNews, querier interfaces.SQLQuerier) ([]*News, error) {
	queryParams := []interface{}{
		filterNews.StartDate,
		filterNews.EndDate,
	}

	query := `
		SELECT 
		id, 
		title, 
		content,
		published_at,
		url
		FROM news
		WHERE published_at BETWEEN ? AND ? 
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

	news := make([]*News, 0)
	for rows.Next() {
		new := News{}
		err := rows.Scan(
			&new.NewsId,
			&new.Title,
			&new.Content,
			&new.PublishedAt,
			&new.Url,
		)
		if err != nil {
			log.Println("Error Scan:", err.Error())
			return nil, err
		}
		news = append(news, &new)
	}

	return news, nil
}
