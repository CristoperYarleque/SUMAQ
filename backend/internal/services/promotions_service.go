package services

import (
	"context"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
)

type Promotion struct {
	PromotionId int
	Description string
	StartDate   string
	EndDate     string
	Url         string
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

type PromotionsServiceInterface interface {
	CreatePromotion(ctx context.Context, bodyPromotion BodyPromotion) error
	GetPromotions(ctx context.Context, filterPromotions FilterPromotions) ([]*Promotion, error)
}

type promotionsService struct {
	promotionsModel models.PromotionsModelInterface
	getConn         func(ctx context.Context) (*interfaces.SQLInterface, error)
}

func NewPromotionsService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) PromotionsServiceInterface {
	return &promotionsService{
		promotionsModel: models.NewPromotionsModel(),
		getConn:         fn,
	}
}

type PromotionsServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.PromotionsModelInterface
}

func NewPromotionsServiceWithModel(arg PromotionsServiceWithModelArguments) PromotionsServiceInterface {
	return &promotionsService{
		promotionsModel: arg.Vw,
		getConn:         arg.Fn,
	}
}

func (c *promotionsService) CreatePromotion(ctx context.Context, bodyPromotion BodyPromotion) error {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	bodyPromotionModel := models.BodyPromotion{
		Description: bodyPromotion.Description,
		StartDate:   bodyPromotion.StartDate,
		EndDate:     bodyPromotion.EndDate,
	}

	err = c.promotionsModel.CreatePromotion(ctx, bodyPromotionModel, db)
	if err != nil {
		return err
	}

	return nil
}

func (c *promotionsService) GetPromotions(ctx context.Context, filterPromotions FilterPromotions) ([]*Promotion, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	filterPromotionsModel := models.FilterPromotions{
		StartDate: filterPromotions.StartDate,
		EndDate:   filterPromotions.EndDate,
	}

	promotions, err := c.promotionsModel.GetPromotions(ctx, filterPromotionsModel, db)
	if err != nil {
		log.Printf("error Model GetPromotions: %v. Request ID: %v\n", err, ctx.Value("requestId"))
		return nil, err
	}

	promotionsResponse := make([]*Promotion, len(promotions))

	for i, promotion := range promotions {
		promotionsResponse[i] = &Promotion{
			PromotionId: promotion.PromotionId,
			Description: promotion.Description,
			StartDate:   promotion.StartDate,
			EndDate:     promotion.EndDate,
			Url:         promotion.Url,
		}
	}

	return promotionsResponse, nil

}
