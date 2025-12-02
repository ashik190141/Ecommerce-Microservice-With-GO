package interfaces

import "Product-Service/dto"

type ProductInterface interface {
	CreateProduct(dto.CreateProductDTO) (bool, string)
	GetProductByID(id int) (dto.GetProductResponse, string)
	UpdateProduct(id int) (bool, string)
	DeleteProduct(id int) (bool, string)
	GetProducts() ([]dto.GetProductResponse, error)
}