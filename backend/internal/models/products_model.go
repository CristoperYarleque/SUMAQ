package models

import (
	"context"
	"fmt"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type BodyProduct struct {
	Name           string
	Description    string
	Price          float64
	Url            string
	CategoryId     int
	EntrepreneurId int
}

type BodyProductUpdate struct {
	Name        string
	Price       float64
	Url         string
	CategoryId  int
	Description string
}

type Products struct {
	ProductId    int
	Name         string
	Price        float64
	Url          string
	Description  string
	CategoryId   int
	CategoryName string
}

type FilterProducts struct {
	CategoryId     int
	EntrepreneurId int
}

type ProductsModelInterface interface {
	CreateProduct(ctx context.Context, bodyProduct BodyProduct, querier interfaces.SQLQuerier) error
	GetProducts(ctx context.Context, filterProducts FilterProducts, querier interfaces.SQLQuerier) ([]*Products, error)
	UpdateProduct(ctx context.Context, productId int, bodyProductUpdate BodyProductUpdate, querier interfaces.SQLQuerier) error
	DeleteProduct(ctx context.Context, productId int, querier interfaces.SQLQuerier) error
	GetProductById(ctx context.Context, productId int, querier interfaces.SQLQuerier) (*Products, error)
}
type productsModel struct {
	Db *interfaces.SQLConnInterface
}

func NewProductsModel() ProductsModelInterface {
	return &productsModel{}
}

func (c *productsModel) CreateProduct(ctx context.Context, bodyProduct BodyProduct, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		bodyProduct.Name,
		bodyProduct.Description,
		bodyProduct.Price,
		bodyProduct.Url,
		bodyProduct.CategoryId,
		bodyProduct.EntrepreneurId,
	}

	query := `
		INSERT INTO products (name, description, price, url, category_id, entrepreneur_id)
		VALUES (?, ?, ?, ?, ?, ?)
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

	fmt.Printf("CreateProduct - Affected lines %d - RequestID %v \n", rowsAffected, ctx.Value("requestId"))
	return nil
}

func (c *productsModel) GetProducts(ctx context.Context, filterProducts FilterProducts, querier interfaces.SQLQuerier) ([]*Products, error) {
	queryParams := []interface{}{
		filterProducts.EntrepreneurId,
	}
	sqlParams := ""

	if filterProducts.CategoryId != 0 {
		queryParams = append(queryParams, filterProducts.CategoryId)
		sqlParams = " AND category_id = ? "
	}

	query := `
		SELECT 
		p.id,
		p.name, 
		p.price, 
		p.url,
		p.description,
		p.category_id,
		c.name
		FROM products p
		INNER JOIN categories c ON c.id = p.category_id
		WHERE entrepreneur_id = ?
		` + sqlParams + `
		ORDER BY 1 DESC;
	`

	stmt, err := querier.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[FATAL]", "Error prepare Query:", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, queryParams...)
	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	products := make([]*Products, 0)
	for rows.Next() {
		product := Products{}
		err := rows.Scan(
			&product.ProductId,
			&product.Name,
			&product.Price,
			&product.Url,
			&product.Description,
			&product.CategoryId,
			&product.CategoryName,
		)
		if err != nil {
			log.Println("Error Scan:", err.Error())
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (c *productsModel) UpdateProduct(ctx context.Context, productId int, bodyProductUpdate BodyProductUpdate, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		bodyProductUpdate.Name,
		bodyProductUpdate.Price,
		bodyProductUpdate.Url,
		bodyProductUpdate.CategoryId,
		bodyProductUpdate.Description,
		productId,
	}

	query := `
		UPDATE products 
		SET name = ?, price = ?, url = ?, category_id = ?, description = ? 
		WHERE id = ?
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

	fmt.Printf("UpdateProduct - Affected lines %d - RequestID %v \n", rowsAffected, ctx.Value("requestId"))
	return nil
}

func (c *productsModel) DeleteProduct(ctx context.Context, productId int, querier interfaces.SQLQuerier) error {
	query := `
		DELETE FROM products 
		WHERE id = ?
	`

	stmt, err := querier.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[FATAL]", "Error prepare Query:", err.Error())
		return err
	}
	defer stmt.Close()

	r, err := stmt.ExecContext(ctx, productId)
	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		log.Printf("Error getting affected rows: %v", err)
		return err
	}

	fmt.Printf("DeleteProduct - Affected lines %d - RequestID %v \n", rowsAffected, ctx.Value("requestId"))
	return nil
}

func (c *productsModel) GetProductById(ctx context.Context, productId int, querier interfaces.SQLQuerier) (*Products, error) {
	queryParams := []interface{}{
		productId,
	}

	query := `
		SELECT 
		id, 
		name, 
		price, 
		url,
		description,
		category_id
		FROM products 
		WHERE id = ?
	`

	row := querier.QueryRowContext(ctx, query, queryParams...)

	product := &Products{}
	err := row.Scan(
		&product.ProductId,
		&product.Name,
		&product.Price,
		&product.Url,
		&product.Description,
		&product.CategoryId,
	)
	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return nil, err
	}

	return product, nil
}
