package service

import (
	helpers "Product-Service/app/response"
	"Product-Service/dto"
	client "Product-Service/handler"
	"Product-Service/interfaces"
	"context"
	"encoding/json"
	"net/http"
)

type productService struct {
	ctx context.Context
	repo interfaces.ProductInterface
	userClient client.UserClient
}

func NewProductService(repo interfaces.ProductInterface) interfaces.ProductService {
	return &productService{ctx: context.Background(), repo: repo, userClient: client.UserClient{}}
}

func (s *productService) CreateProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[dto.GetProductResponse] {
	var newProduct dto.CreateProductDTO
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusBadRequest, "Invalid request body", dto.GetProductResponse{})
	}
	
	isUserExist, _ := s.userClient.IsUserExist(s.ctx, newProduct.Email)
	if(!isUserExist.Success) {
		return *helpers.StandardApiResponse(false, http.StatusBadRequest, "User does not exist", dto.GetProductResponse{})
	}
	
	created := repo.CreateProduct(newProduct)
	if(!created) {
		return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Failed to create product", dto.GetProductResponse{})
	}

	return helpers.ApiResponse[dto.GetProductResponse]{}
}

func (s *productService) GetByIDProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[dto.GetProductResponse] {
	return helpers.ApiResponse[dto.GetProductResponse]{}
}

func (s *productService) UpdateProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[dto.GetProductResponse] {
	return helpers.ApiResponse[dto.GetProductResponse]{}
}

func (s *productService) DeleteProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[dto.GetProductResponse] {
	return helpers.ApiResponse[dto.GetProductResponse]{}
}

func (s *productService) GetProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[dto.GetProductResponse] {
	return helpers.ApiResponse[dto.GetProductResponse]{}
}