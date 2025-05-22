package services

import (
	"context"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
)

type BodyInfoEntrepreneur struct {
	EntrepreneurId int
	Type           string
	Url            string
	Description    string
}

type InfoEntrepreneurs struct {
	Id          int
	Url         string
	Description string
}

type FilterInfoEntrepreneurs struct {
	EntrepreneurId int
	Type           string
}

type InfoEntrepreneursServiceInterface interface {
	CreateInfoEntrepreneur(ctx context.Context, bodyInfoEntrepreneur BodyInfoEntrepreneur) error
	GetInfoEntrepreneurs(ctx context.Context, filterInfoEntrepreneurs FilterInfoEntrepreneurs) (*InfoEntrepreneurs, error)
}

type infoEntrepreneursService struct {
	infoEntrepreneursModel models.InfoEntrepreneursModelInterface
	getConn                func(ctx context.Context) (*interfaces.SQLInterface, error)
}

func NewInfoEntrepreneursService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) InfoEntrepreneursServiceInterface {
	return &infoEntrepreneursService{
		infoEntrepreneursModel: models.NewInfoEntrepreneursModel(),
		getConn:                fn,
	}
}

type InfoEntrepreneursServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.InfoEntrepreneursModelInterface
}

func NewInfoEntrepreneursServiceWithModel(arg InfoEntrepreneursServiceWithModelArguments) InfoEntrepreneursServiceInterface {
	return &infoEntrepreneursService{
		infoEntrepreneursModel: arg.Vw,
		getConn:                arg.Fn,
	}
}

func (c *infoEntrepreneursService) CreateInfoEntrepreneur(ctx context.Context, bodyInfoEntrepreneur BodyInfoEntrepreneur) error {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	bodyInfoEntrepreneurModel := models.BodyInfoEntrepreneur{
		EntrepreneurId: bodyInfoEntrepreneur.EntrepreneurId,
		Type:           bodyInfoEntrepreneur.Type,
		Url:            bodyInfoEntrepreneur.Url,
		Description:    bodyInfoEntrepreneur.Description,
	}

	err = c.infoEntrepreneursModel.CreateInfoEntrepreneur(ctx, bodyInfoEntrepreneurModel, db)
	if err != nil {
		return err
	}

	return nil
}

func (c *infoEntrepreneursService) GetInfoEntrepreneurs(ctx context.Context, filterInfoEntrepreneurs FilterInfoEntrepreneurs) (*InfoEntrepreneurs, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	filterInfoEntrepreneursModel := models.FilterInfoEntrepreneurs{
		EntrepreneurId: filterInfoEntrepreneurs.EntrepreneurId,
		Type:           filterInfoEntrepreneurs.Type,
	}

	infoEntrepreneurs, err := c.infoEntrepreneursModel.GetInfoEntrepreneurs(ctx, filterInfoEntrepreneursModel, db)
	if err != nil {
		return nil, err
	}

	infoEntrepreneursResponse := InfoEntrepreneurs{
		Id:          infoEntrepreneurs.Id,
		Url:         infoEntrepreneurs.Url,
		Description: infoEntrepreneurs.Description,
	}

	return &infoEntrepreneursResponse, nil
}
