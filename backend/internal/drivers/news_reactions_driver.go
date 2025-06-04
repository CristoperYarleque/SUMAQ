package drivers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/services"
)

type NewsReactionsDriverInterface interface {
	GetReactionInfo(w http.ResponseWriter, r *http.Request)
	UpsertReaction(w http.ResponseWriter, r *http.Request)
	DeleteReaction(w http.ResponseWriter, r *http.Request)
}

type newsReactionsDriver struct {
	NewsReactionsService services.NewsReactionsServiceInterface
	BaseDriverInterface
}

func NewNewsReactionsDriver(db database.Connections) NewsReactionsDriverInterface {
	return &newsReactionsDriver{
		NewsReactionsService: services.NewsReactionsService(func(ctx context.Context) (*interfaces.SQLInterface, error) {
			return &db.SQL, nil
		}),
		BaseDriverInterface: &baseDriver{},
	}
}

func NewNewsReactionsDriverWithServices(se services.NewsReactionsServiceInterface) NewsReactionsDriverInterface {
	return &newsReactionsDriver{
		NewsReactionsService: se,
		BaseDriverInterface:  &baseDriver{},
	}
}

func (c *newsReactionsDriver) GetReactionInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	newsIDStr := r.URL.Query().Get("newsId")
	userIDStr := r.URL.Query().Get("userId")

	if newsIDStr == "" && userIDStr == "" {
		c.handleError(w, errors.New("newsId and userId are required"), http.StatusBadRequest)
		return
	}

	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		c.handleError(w, errors.New("newsId must be a valid integer"), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.handleError(w, errors.New("userId must be a valid integer"), http.StatusBadRequest)
		return
	}

	filter := services.BodyUpsertReaction{
		UserID: userID,
		NewsID: newsID,
	}

	reactionInfo, err := c.NewsReactionsService.GetReactionInfo(ctx, filter)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, reactionInfo, http.StatusOK)
}

func (c *newsReactionsDriver) UpsertReaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body services.BodyUpsertReaction
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		c.handleError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
		return
	}

	if body.NewsID == 0 || body.UserID == 0 || body.Reaction == "" {
		c.handleError(w, errors.New("all fields (news_id, user_id, reaction) are required"), http.StatusBadRequest)
		return
	}

	if err := c.NewsReactionsService.UpsertReaction(ctx, body); err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusCreated)
}

func (c *newsReactionsDriver) DeleteReaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	newsIDStr := r.URL.Query().Get("newsId")
	userIDStr := r.URL.Query().Get("userId")

	if newsIDStr == "" && userIDStr == "" {
		c.handleError(w, errors.New("newsId and userId are required"), http.StatusBadRequest)
		return
	}

	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		c.handleError(w, errors.New("newsId must be a valid integer"), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.handleError(w, errors.New("userId must be a valid integer"), http.StatusBadRequest)
		return
	}

	filter := services.BodyUpsertReaction{
		UserID: userID,
		NewsID: newsID,
	}

	err = c.NewsReactionsService.DeleteProduct(ctx, filter)
	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleSuccess(w, nil, http.StatusOK)
}
