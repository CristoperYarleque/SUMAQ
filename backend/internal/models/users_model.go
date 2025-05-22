package models

import (
	"context"
	"fmt"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type Users struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
}

type BodyUser struct {
	Name     string
	Email    string
	Password string
	Role     string
	Url      string
}

type FilterUser struct {
	Email string
	Role  string
}

type UsersModelInterface interface {
	CreateUser(ctx context.Context, bodyUser BodyUser, querier interfaces.SQLQuerier) error
	GetUser(ctx context.Context, filterUser FilterUser, querier interfaces.SQLQuerier) (*Users, error)
}
type usersModel struct {
	Db *interfaces.SQLConnInterface
}

func NewUsersModel() UsersModelInterface {
	return &usersModel{}
}

func (c *usersModel) CreateUser(ctx context.Context, bodyUser BodyUser, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		bodyUser.Name,
		bodyUser.Email,
		bodyUser.Password,
		bodyUser.Role,
		bodyUser.Url,
	}

	fmt.Println("queryParams", queryParams)

	query := `
		INSERT INTO users (name, email, password, role, url)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := querier.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[FATAL]", "Error prepare Query:", err.Error())
		return err
	}
	defer stmt.Close()

	r, err := stmt.ExecContext(ctx, queryParams...)
	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		log.Printf("Error getting affected rows: %v", err)
		return err
	}

	fmt.Printf("CreateUser - Affected lines %d - RequestID %v \n", rowsAffected, ctx.Value("requestId"))
	return nil
}

func (c *usersModel) GetUser(ctx context.Context, filterUser FilterUser, querier interfaces.SQLQuerier) (*Users, error) {
	queryParams := []interface{}{
		filterUser.Email,
		filterUser.Role,
	}

	query := `
		SELECT 
		id, 
		name, 
		email, 
		password, 
		role 
		FROM users 
		WHERE email = ? AND role = ?
	`

	rows := querier.QueryRowContext(ctx, query, queryParams...)

	user := &Users{}
	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return nil, err
	}

	return user, nil
}
