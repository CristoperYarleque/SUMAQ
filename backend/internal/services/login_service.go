package services

import (
	"context"
	"errors"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/models"
	"github.com/cyarleque/sumaq/internal/utils"
)

type Login struct {
	UserId int
	Name   string
	Role   string
	Url    string
	Token  string
}

type BodyLogin struct {
	Email    string
	Password string
	Role     string
}

type LoginServiceInterface interface {
	LoginUser(ctx context.Context, bodyLogin BodyLogin) (*Login, error)
}

type loginService struct {
	usersModel models.UsersModelInterface
	getConn    func(ctx context.Context) (*interfaces.SQLInterface, error)
	utils      utils.UtilsInterface
}

func NewLoginService(fn func(ctx context.Context) (*interfaces.SQLInterface, error)) LoginServiceInterface {
	return &loginService{
		usersModel: models.NewUsersModel(),
		getConn:    fn,
		utils:      &utils.Utils{},
	}
}

type LoginServiceWithModelArguments struct {
	Fn func(ctx context.Context) (*interfaces.SQLInterface, error)
	Vw models.UsersModelInterface
	Uw utils.UtilsInterface
}

func NewLoginServiceWithModel(arg LoginServiceWithModelArguments) LoginServiceInterface {
	return &loginService{
		usersModel: arg.Vw,
		getConn:    arg.Fn,
		utils:      arg.Uw,
	}
}

func (c *loginService) LoginUser(ctx context.Context, bodyLogin BodyLogin) (*Login, error) {
	dbPtr, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	db := *dbPtr

	bodyLoginModel := models.FilterUser{
		Email: bodyLogin.Email,
		Role:  bodyLogin.Role,
	}

	user, err := c.usersModel.GetUser(ctx, bodyLoginModel, db)
	if err != nil {
		return nil, err
	}

	err = c.utils.CheckPassword(user.Password, bodyLogin.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := c.utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	login := Login{
		UserId: user.ID,
		Name:   user.Name,
		Role:   user.Role,
		Url:    user.Url,
		Token:  token,
	}

	return &login, nil
}
