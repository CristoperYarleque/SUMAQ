package services

import (
	"context"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
)

type News struct {
	NewsId      int
	Title       string
	Content     string
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

type NewsServiceInterface interface {
	CreateNews(ctx context.Context, bodyNews BodyNews) error
	GetNews(ctx context.Context, filterNews FilterNews) ([]*News, error)
}

type newsService struct {
	newsModel models.NewsModelInterface
	getConn   func(ctx context.Context) (*interfaces.SQLInterface, error)
}

func NewNewsService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) NewsServiceInterface {
	return &newsService{
		newsModel: models.NewNewsModel(),
		getConn:   fn,
	}
}

type NewsServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.NewsModelInterface
}

func NewNewsServiceWithModel(arg NewsServiceWithModelArguments) NewsServiceInterface {
	return &newsService{
		newsModel: arg.Vw,
		getConn:   arg.Fn,
	}
}

func (c *newsService) CreateNews(ctx context.Context, bodyNews BodyNews) error {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	bodyNewsModel := models.BodyNews{
		Title:       bodyNews.Title,
		Content:     bodyNews.Content,
		PublishedAt: bodyNews.PublishedAt,
	}

	err = c.newsModel.CreateNews(ctx, bodyNewsModel, db)
	if err != nil {
		return err
	}

	return nil
}

func (c *newsService) GetNews(ctx context.Context, filterNews FilterNews) ([]*News, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	filterNewsModel := models.FilterNews{
		StartDate: filterNews.StartDate,
		EndDate:   filterNews.EndDate,
	}

	news, err := c.newsModel.GetNews(ctx, filterNewsModel, db)
	if err != nil {
		log.Printf("error Model GetNews: %v. Request ID: %v\n", err, ctx.Value("requestId"))
		return nil, err
	}

	newsResponse := make([]*News, len(news))

	for i, news := range news {
		newsResponse[i] = &News{
			NewsId:      news.NewsId,
			Title:       news.Title,
			Content:     news.Content,
			PublishedAt: news.PublishedAt,
		}
	}

	return newsResponse, nil

}
