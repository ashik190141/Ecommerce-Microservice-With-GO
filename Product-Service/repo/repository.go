package repo

import (
	"Product-Service/dto"
	"Product-Service/interfaces"
	"context"
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
	return false
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