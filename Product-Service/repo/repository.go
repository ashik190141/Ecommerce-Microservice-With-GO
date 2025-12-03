package repo

import (
	"Product-Service/dto"
	"Product-Service/interfaces"
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewProductRepository(db *sqlx.DB) interfaces.ProductInterface {
	return &ProductRepository{
		db:  db,
		ctx: context.Background(),
	}
}

func (r *ProductRepository) CreateProduct(product dto.CreateProductDTO) dto.GetProductResponse {
	query := `
        INSERT INTO products (name, price, stock, useremail, sku)
        VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, price, stock, useremail, sku, created_at, updated_at
    `

	var createdProduct dto.GetProductResponse

	err := r.db.QueryRowContext(r.ctx, query, product.Name, product.Price, product.Stock, product.Email, product.Sku).Scan(&createdProduct.Id, &createdProduct.Name, &createdProduct.Price, &createdProduct.Stock, &createdProduct.Email, &createdProduct.Sku, &createdProduct.CreatedAt, &createdProduct.UpdatedAt)

	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return dto.GetProductResponse{}
		}
		return dto.GetProductResponse{}
	}

	return createdProduct
}

func (r *ProductRepository) GetProductByID(id int) (dto.GetProductResponse, string) {
	return dto.GetProductResponse{}, "not implemented"
}

func (r *ProductRepository) UpdateProduct(id int, product dto.CreateProductDTO) bool {
	return false
}

func (r *ProductRepository) DeleteProduct(id int) bool {
	return false
}

func (r *ProductRepository) GetProducts() ([]dto.GetProductResponse, error) {
	var products []dto.GetProductResponse
	query := `SELECT id, name, price, stock, useremail, sku, created_at, updated_at FROM products`
	err := r.db.SelectContext(r.ctx, &products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}