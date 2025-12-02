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

func (r *ProductRepository) CreateProduct(product dto.CreateProductDTO) (bool) {
	query := `
        INSERT INTO products (name, price, stock, userEmail, sku)
        VALUES ($1, $2, $3, $4, $5)
    `
	if err := r.db.QueryRowContext(r.ctx, query, product.Name, product.Price, product.Stock, product.Email, product.Sku).Scan(&product.Name, &product.Price, &product.Stock, &product.Email, &product.Sku); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return false
		}
		return false
	}
	return true
}

func (r *ProductRepository) GetProductByID(id int) (dto.GetProductResponse, string) {
	return dto.GetProductResponse{}, "not implemented"
}

func (r *ProductRepository) UpdateProduct(id int) bool {
	return false
}

func (r *ProductRepository) DeleteProduct(id int) bool {
	return false
}

func (r *ProductRepository) GetProducts() ([]dto.GetProductResponse, error) {
	return nil, nil
}