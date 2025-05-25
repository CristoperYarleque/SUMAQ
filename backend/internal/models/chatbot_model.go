package models

import (
	"context"
	"fmt"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type Chatbot struct {
	ChatbotId int
	Question  string
	Answer    string
}

type BodyChatbot struct {
	Question string
	Answer   string
}

type ChatbotModelInterface interface {
	CreateChatbot(ctx context.Context, bodyChatbot BodyChatbot, querier interfaces.SQLQuerier) error
	GetChatbot(ctx context.Context, querier interfaces.SQLQuerier) ([]*Chatbot, error)
}
type chatbotModel struct {
	Db *interfaces.SQLConnInterface
}

func NewChatbotModel() ChatbotModelInterface {
	return &chatbotModel{}
}

func (c *chatbotModel) CreateChatbot(ctx context.Context, bodyChatbot BodyChatbot, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		bodyChatbot.Question,
		bodyChatbot.Answer,
	}

	query := `
		INSERT INTO question_answer (question, answer)
		VALUES (?, ?)
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

	fmt.Printf("CreateChatbot - Affected lines %d - RequestID %v \n", rowsAffected, ctx.Value("requestId"))
	return nil
}

func (c *chatbotModel) GetChatbot(ctx context.Context, querier interfaces.SQLQuerier) ([]*Chatbot, error) {
	queryParams := []interface{}{}

	query := `
		SELECT 
		id, 
		question, 
		answer
		FROM question_answer
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

	chatbots := make([]*Chatbot, 0)
	for rows.Next() {
		chatbot := Chatbot{}
		err := rows.Scan(
			&chatbot.ChatbotId,
			&chatbot.Question,
			&chatbot.Answer,
		)
		if err != nil {
			log.Println("Error Scan:", err.Error())
			return nil, err
		}
		chatbots = append(chatbots, &chatbot)
	}

	return chatbots, nil
}
