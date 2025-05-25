package services

import (
	"context"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
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

type ChatbotServiceInterface interface {
	CreateChatbot(ctx context.Context, bodyChatbot BodyChatbot) error
	GetChatbot(ctx context.Context) ([]*Chatbot, error)
}

type chatbotService struct {
	chatbotModel models.ChatbotModelInterface
	getConn      func(ctx context.Context) (*interfaces.SQLInterface, error)
}

func NewChatbotService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) ChatbotServiceInterface {
	return &chatbotService{
		chatbotModel: models.NewChatbotModel(),
		getConn:      fn,
	}
}

type ChatbotServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.ChatbotModelInterface
}

func NewChatbotServiceWithModel(arg ChatbotServiceWithModelArguments) ChatbotServiceInterface {
	return &chatbotService{
		chatbotModel: arg.Vw,
		getConn:      arg.Fn,
	}
}

func (c *chatbotService) CreateChatbot(ctx context.Context, bodyChatbot BodyChatbot) error {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	bodyChatbotModel := models.BodyChatbot{
		Question: bodyChatbot.Question,
		Answer:   bodyChatbot.Answer,
	}

	err = c.chatbotModel.CreateChatbot(ctx, bodyChatbotModel, db)
	if err != nil {
		return err
	}

	return nil
}

func (c *chatbotService) GetChatbot(ctx context.Context) ([]*Chatbot, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	chatbot, err := c.chatbotModel.GetChatbot(ctx, db)
	if err != nil {
		log.Printf("error Model GetChatbot: %v. Request ID: %v\n", err, ctx.Value("requestId"))
		return nil, err
	}

	chatbotResponse := make([]*Chatbot, len(chatbot))

	for i, chatbot := range chatbot {
		chatbotResponse[i] = &Chatbot{
			ChatbotId: chatbot.ChatbotId,
			Question:  chatbot.Question,
			Answer:    chatbot.Answer,
		}
	}

	return chatbotResponse, nil

}
