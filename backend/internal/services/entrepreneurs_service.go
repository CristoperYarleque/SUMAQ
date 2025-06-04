package services

import (
	"context"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
)

type Entrepreneurs struct {
	Id    int
	Name  string
	Email string
	Url   string
	Role  string
}

type BodyEntrepreneurUpdate struct {
	Url string
}

type FilterEntrepreneurs struct {
	CategoryId     int
	EntrepreneurId int
}

type EntrepreneursServiceInterface interface {
	GetEntrepreneurs(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs) ([]*Entrepreneurs, error)
	UpdateUrlEntrepreneur(ctx context.Context, entrepreneurId int, bodyEntrepreneurUpdate BodyEntrepreneurUpdate) error
	GetEntrepreneur(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs) (*Entrepreneurs, error)
}

type entrepreneursService struct {
	entrepreneursModel models.EntrepreneursModelInterface
	getConn            func(ctx context.Context) (*interfaces.SQLInterface, error)
}

func NewEntrepreneursService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) EntrepreneursServiceInterface {
	return &entrepreneursService{
		entrepreneursModel: models.NewEntrepreneursModel(),
		getConn:            fn,
	}
}

type EntrepreneursServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.EntrepreneursModelInterface
}

func NewEntrepreneursServiceWithModel(arg EntrepreneursServiceWithModelArguments) EntrepreneursServiceInterface {
	return &entrepreneursService{
		entrepreneursModel: arg.Vw,
		getConn:            arg.Fn,
	}
}

func (c *entrepreneursService) GetEntrepreneurs(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs) ([]*Entrepreneurs, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	filterEntrepreneursModel := models.FilterEntrepreneurs{
		CategoryId: filterEntrepreneurs.CategoryId,
	}

	entrepreneurs, err := c.entrepreneursModel.GetEntrepreneurs(ctx, filterEntrepreneursModel, db)
	if err != nil {
		return nil, err
	}

	entrepreneursResponse := make([]*Entrepreneurs, len(entrepreneurs))

	for i, entrepreneur := range entrepreneurs {
		entrepreneursResponse[i] = &Entrepreneurs{
			Id:    entrepreneur.Id,
			Name:  entrepreneur.Name,
			Email: entrepreneur.Email,
			Url:   entrepreneur.Url,
			Role:  entrepreneur.Role,
		}
	}

	return entrepreneursResponse, nil
}

func (c *entrepreneursService) GetEntrepreneur(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs) (*Entrepreneurs, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	filterEntrepreneursModel := models.FilterEntrepreneurs{
		EntrepreneurId: filterEntrepreneurs.EntrepreneurId,
	}

	entrepreneur, err := c.entrepreneursModel.GetEntrepreneur(ctx, filterEntrepreneursModel, db)
	if err != nil {
		return nil, err
	}
	if entrepreneur == nil {
		return nil, nil
	}
	return &Entrepreneurs{
		Id:    entrepreneur.Id,
		Name:  entrepreneur.Name,
		Email: entrepreneur.Email,
		Url:   entrepreneur.Url,
		Role:  entrepreneur.Role,
	}, nil
}

func (c *entrepreneursService) UpdateUrlEntrepreneur(ctx context.Context, entrepreneurId int, bodyEntrepreneurUpdate BodyEntrepreneurUpdate) error {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	bodyEntrepreneurUpdateModel := models.BodyEntrepreneurUpdate{
		Url: bodyEntrepreneurUpdate.Url,
	}

	err = c.entrepreneursModel.UpdateUrlEntrepreneur(ctx, entrepreneurId, bodyEntrepreneurUpdateModel, db)
	if err != nil {
		return err
	}

	return nil
}
