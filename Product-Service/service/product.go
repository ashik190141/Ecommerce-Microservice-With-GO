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
	ctx        context.Context
	repo       interfaces.ProductInterface
	userClient *client.UserClient
}

func NewProductService(repo interfaces.ProductInterface, uc *client.UserClient) interfaces.ProductService {
	return &productService{ctx: context.Background(), repo: repo, userClient: uc}
}

func (s *productService) CreateProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[dto.GetProductResponse] {
	var newProduct dto.CreateProductDTO
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusBadRequest, "Invalid request body", dto.GetProductResponse{})
	}

	isUserExist, err := s.userClient.IsUserExist(s.ctx, newProduct.Email)
	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Failed to validate user", dto.GetProductResponse{})
	}
	if !isUserExist.Success {
		return *helpers.StandardApiResponse(false, http.StatusBadRequest, "User does not exist", dto.GetProductResponse{})
	}

	created := repo.CreateProduct(newProduct)
	if (created == dto.GetProductResponse{}) {
		return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Failed to create product", dto.GetProductResponse{})
	}

	return *helpers.StandardApiResponse(false, http.StatusOK, "Create Product Successfully", created)
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

func (s *productService) GetProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[[]dto.GetProductResponse] {
	products, err := repo.GetProducts()
	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Failed to get products", []dto.GetProductResponse{})
	}
	return *helpers.StandardApiResponse(true, http.StatusOK, "Products retrieved successfully", products)
}
