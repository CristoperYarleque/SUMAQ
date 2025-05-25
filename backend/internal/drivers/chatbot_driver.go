package drivers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/services"
)

type ChatbotDriverInterface interface {
	CreateChatbot(w http.ResponseWriter, r *http.Request)
	GetChatbot(w http.ResponseWriter, r *http.Request)
}
type chatbotDriver struct {
	ChatbotService services.ChatbotServiceInterface
	BaseDriverInterface
}

func NewChatbotDriver(db database.Connections) ChatbotDriverInterface {
	return &chatbotDriver{
		ChatbotService: services.NewChatbotService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func NewChatbotDriverWithServices(se services.ChatbotServiceInterface) ChatbotDriverInterface {
	return &chatbotDriver{
		ChatbotService:      se,
		BaseDriverInterface: &baseDriver{},
	}
}

func (c *chatbotDriver) CreateChatbot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var bodyChatbot services.BodyChatbot
	if err := json.NewDecoder(r.Body).Decode(&bodyChatbot); err != nil {
		c.handleError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
		return
	}

	// Validaciones básicas
	if bodyChatbot.Question == "" || bodyChatbot.Answer == "" {
		c.handleError(w, errors.New("all fields (question, answer) are required"), http.StatusBadRequest)
		return
	}

	err := c.ChatbotService.CreateChatbot(ctx, bodyChatbot)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusCreated)
}

func (c *chatbotDriver) GetChatbot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	chatbot, err := c.ChatbotService.GetChatbot(ctx)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, chatbot, http.StatusOK)
}
