package repo

import (
	"Product-Service/dto"
	"Product-Service/interfaces"
	"context"
	"fmt"
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

func (r *ProductRepository) GetProductByID(id int) dto.GetProductResponse {
	var products dto.GetProductResponse
	query := `SELECT id, name, price, stock, useremail, sku, created_at, updated_at FROM products WHERE id=$1`
	err := r.db.GetContext(r.ctx, &products, query, id)
	if err != nil {
		return dto.GetProductResponse{}
	}
	return products
}

func (r *ProductRepository) UpdateProduct(id int, product dto.UpdateProductDTO) dto.GetProductResponse {
	var setParts []string
	var args []interface{}
	argIndex := 1

	if product.Name != "" {
		setParts = append(setParts, "name = $"+fmt.Sprintf("%d", argIndex))
		args = append(args, product.Name)
		argIndex++
	}
	if product.Price != 0 {
		setParts = append(setParts, "price = $"+fmt.Sprintf("%d", argIndex))
		args = append(args, product.Price)
		argIndex++
	}
	if product.Stock != 0 {
		setParts = append(setParts, "stock = $"+fmt.Sprintf("%d", argIndex))
		args = append(args, product.Stock)
		argIndex++
	}
	if product.Sku != "" {
		setParts = append(setParts, "sku = $"+fmt.Sprintf("%d", argIndex))
		args = append(args, product.Sku)
		argIndex++
	}

	if len(setParts) == 0 {
		return dto.GetProductResponse{}
	}

	setParts = append(setParts, "updated_at = NOW()")
	args = append(args, id)

	query := "UPDATE products SET " + strings.Join(setParts, ", ") + " WHERE id = $" + fmt.Sprintf("%d", argIndex) + " RETURNING id, name, price, stock, useremail, sku, created_at, updated_at"

	var updatedProduct dto.GetProductResponse
	err := r.db.QueryRowContext(r.ctx, query, args...).Scan(
		&updatedProduct.Id,
		&updatedProduct.Name,
		&updatedProduct.Price,
		&updatedProduct.Stock,
		&updatedProduct.Email,
		&updatedProduct.Sku,
		&updatedProduct.CreatedAt,
		&updatedProduct.UpdatedAt,
	)

	if err != nil {
		return dto.GetProductResponse{}
	}

	return updatedProduct
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