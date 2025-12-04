package interfaces

import (
	helpers "Product-Service/app/response"
	"Product-Service/dto"
	"net/http"
	"time"
)

type ProductInterface interface {
	CreateProduct(product dto.CreateProductDTO) dto.GetProductResponse
	GetProductByID(id int) (dto.GetProductResponse)
	UpdateProduct(id int, product dto.UpdateProductDTO) dto.GetProductResponse
	DeleteProduct(id int) bool
	GetProducts() ([]dto.GetProductResponse, error)
}

type ProductService interface {
	CreateProductService(r *http.Request, repo ProductInterface) (helpers.ApiResponse[dto.GetProductResponse])
	GetByIDProductService(r *http.Request, repo ProductInterface) (helpers.ApiResponse[dto.GetProductResponse])
	UpdateProductService(r *http.Request, repo ProductInterface) (helpers.ApiResponse[dto.GetProductResponse])
	DeleteProductService(r *http.Request, repo ProductInterface) (helpers.ApiResponse[dto.GetProductResponse])
	GetProductService(r *http.Request, repo ProductInterface) (helpers.ApiResponse[[]dto.GetProductResponse])
}

type ProductHandler interface {
	CreateProductHandler(w http.ResponseWriter, r *http.Request)
	GetProductByIDHandler(w http.ResponseWriter, r *http.Request)
	UpdateProductHandler(w http.ResponseWriter, r *http.Request)
	// DeleteProductHandler(w http.ResponseWriter, r *http.Request)
	GetProductsHandler(w http.ResponseWriter, r *http.Request)
}

type ProductRedis interface {
	GetProductFromCache(productID string) ([]dto.GetProductResponse, error)
	SetProductToCache(key string, product dto.GetProductResponse) bool
	SetExpireTimeFromCache(key string, duration time.Duration)
	GetProductByIdFromCache(key string, id string) dto.GetProductResponse
}