package interfaces

import "Product-Service/dto"

type ProductInterface interface {
	CreateProduct(dto.CreateProductDTO) bool
	GetProductByID(id int) (dto.GetProductResponse, string)
	UpdateProduct(id int) bool
	DeleteProduct(id int) bool
	GetProducts() ([]dto.GetProductResponse, error)
}