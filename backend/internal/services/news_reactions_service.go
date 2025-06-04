package services

import (
	"context"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
)

type BodyUpsertReaction struct {
	UserID   int
	NewsID   int
	Reaction string
}

type NewsReactionInfo struct {
	UserReaction string
	LikeCount    int
	LoveCount    int
}

type NewsReactionsServiceInterface interface {
	UpsertReaction(ctx context.Context, body BodyUpsertReaction) error
	GetReactionInfo(ctx context.Context, filter BodyUpsertReaction) (*NewsReactionInfo, error)
	DeleteProduct(ctx context.Context, filter BodyUpsertReaction) error
}

type newsReactionsService struct {
	newsReactionsModel models.NewsReactionsModelInterface
	getConn            func(ctx context.Context) (*interfaces.SQLInterface, error)
}

func NewsReactionsService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) NewsReactionsServiceInterface {
	return &newsReactionsService{
		newsReactionsModel: models.NewNewsReactionsModel(),
		getConn:            fn,
	}
}

type NewsReactionsServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.NewsReactionsModelInterface
}

func NewNewsReactionsServiceWithModel(arg NewsReactionsServiceWithModelArguments) NewsReactionsServiceInterface {
	return &newsReactionsService{
		newsReactionsModel: arg.Vw,
		getConn:            arg.Fn,
	}
}

func (s *newsReactionsService) UpsertReaction(ctx context.Context, body BodyUpsertReaction) error {
	dbPtr, err := s.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	modelBody := models.FiltersNewsReactions{
		UserID:   body.UserID,
		NewsID:   body.NewsID,
		Reaction: body.Reaction,
	}

	return s.newsReactionsModel.UpsertReaction(ctx, modelBody, db)
}

func (s *newsReactionsService) GetReactionInfo(ctx context.Context, filter BodyUpsertReaction) (*NewsReactionInfo, error) {
	dbPtr, err := s.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	filterNewsReactions := models.FiltersNewsReactions{
		UserID:   filter.UserID,
		NewsID:   filter.NewsID,
		Reaction: filter.Reaction,
	}

	modelInfo, err := s.newsReactionsModel.GetReactionSummary(ctx, filterNewsReactions, db)
	if err != nil {
		return nil, err
	}

	return &NewsReactionInfo{
		UserReaction: modelInfo.UserReaction,
		LikeCount:    modelInfo.LikeCount,
		LoveCount:    modelInfo.LoveCount,
	}, nil
}

func (s *newsReactionsService) DeleteProduct(ctx context.Context, filter BodyUpsertReaction) error {
	dbPtr, err := s.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	filterDelete := models.FiltersNewsReactions{
		UserID: filter.UserID,
		NewsID: filter.NewsID,
	}

	err = s.newsReactionsModel.DeleteReactions(ctx, filterDelete, db)
	if err != nil {
		return err
	}

	return nil
}
