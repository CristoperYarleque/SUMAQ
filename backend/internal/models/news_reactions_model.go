package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type NewsReactionSummary struct {
	UserReaction string
	LikeCount    int
	LoveCount    int
}

type FiltersNewsReactions struct {
	UserID   int
	NewsID   int
	Reaction string
}

type NewsReactionsModelInterface interface {
	UpsertReaction(ctx context.Context, filters FiltersNewsReactions, querier interfaces.SQLQuerier) error
	GetReactionSummary(ctx context.Context, filters FiltersNewsReactions, querier interfaces.SQLQuerier) (*NewsReactionSummary, error)
	DeleteReactions(ctx context.Context, filters FiltersNewsReactions, querier interfaces.SQLQuerier) error
}

type newsReactionsModel struct {
	Db *interfaces.SQLConnInterface
}

func NewNewsReactionsModel() NewsReactionsModelInterface {
	return &newsReactionsModel{}
}

func (m *newsReactionsModel) UpsertReaction(ctx context.Context, filters FiltersNewsReactions, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		filters.UserID,
		filters.NewsID,
		filters.Reaction,
	}

	query := `
		INSERT INTO news_reactions (user_id, new_id, reaction)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE reaction = VALUES(reaction)
	`

	stmt, err := querier.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[FATAL] Prepare UpsertReaction:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, queryParams...)
	if err != nil {
		log.Println("[FATAL] Exec UpsertReaction:", err)
		return err
	}

	fmt.Printf("UpsertReaction - user %d, news %d, reaction %s - RequestID %v\n", filters.UserID, filters.NewsID, filters.Reaction, ctx.Value("requestId"))
	return nil
}

func (m *newsReactionsModel) GetReactionSummary(ctx context.Context, filters FiltersNewsReactions, querier interfaces.SQLQuerier) (*NewsReactionSummary, error) {
	var summary NewsReactionSummary

	queryParams := []interface{}{
		filters.UserID,
		filters.NewsID,
		filters.NewsID,
		filters.NewsID,
	}

	query := `
		SELECT
		COALESCE(
			(SELECT reaction FROM news_reactions WHERE user_id = ? AND new_id = ? LIMIT 1),
			''
		) AS user_reaction,
		COALESCE(
			(SELECT COUNT(*) FROM news_reactions WHERE new_id = ? AND reaction = 'like'),
			0
		) AS like_count,
		COALESCE(
			(SELECT COUNT(*) FROM news_reactions WHERE new_id = ? AND reaction = 'love'),
			0
		) AS love_count;
	`

	row := querier.QueryRowContext(ctx, query, queryParams...)

	err := row.Scan(
		&summary.UserReaction,
		&summary.LikeCount,
		&summary.LoveCount,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[FATAL] GetReactionSummary error: %v", err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		return &summary, nil
	}

	return &summary, nil
}

func (m *newsReactionsModel) DeleteReactions(ctx context.Context, filters FiltersNewsReactions, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		filters.UserID,
		filters.NewsID,
	}

	query := `
		DELETE FROM news_reactions
		WHERE user_id = ? AND new_id = ?
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

	fmt.Printf("DeleteReactions - Affected lines %d - user %d, news %d - RequestID %v\n", rowsAffected, filters.UserID, filters.NewsID, ctx.Value("requestId"))
	return nil
}
