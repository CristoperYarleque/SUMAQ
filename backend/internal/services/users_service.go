package services

import (
	"context"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
	"github.com/cyarleque/sumaq/internal/utils"
)

type BodyUser struct {
	Name     string
	Email    string
	Password string
	Role     string
	Url      string
}

type UsersServiceInterface interface {
	CreateUser(ctx context.Context, bodyUser BodyUser) error
}

type usersService struct {
	usersModel models.UsersModelInterface
	getConn    func(ctx context.Context) (*interfaces.SQLInterface, error)
	utils      utils.UtilsInterface
}

func NewUsersService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) UsersServiceInterface {
	return &usersService{
		usersModel: models.NewUsersModel(),
		getConn:    fn,
		utils:      &utils.Utils{},
	}
}

type UsersServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.UsersModelInterface
	Uw utils.UtilsInterface
}

func NewUsersServiceWithModel(arg UsersServiceWithModelArguments) UsersServiceInterface {
	return &usersService{
		usersModel: arg.Vw,
		getConn:    arg.Fn,
		utils:      arg.Uw,
	}
}

func (c *usersService) CreateUser(ctx context.Context, bodyUser BodyUser) error {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return err
	}
	db := *dbPtr

	hashedPassword, err := c.utils.HashPassword(bodyUser.Password)
	if err != nil {
		return err
	}

	bodyUserModel := models.BodyUser{
		Name:     bodyUser.Name,
		Email:    bodyUser.Email,
		Password: hashedPassword,
		Role:     bodyUser.Role,
		Url:      bodyUser.Url,
	}

	err = c.usersModel.CreateUser(ctx, bodyUserModel, db)
	if err != nil {
		return err
	}

	return nil
}
