package service

import (
	helpers "Product-Service/app/response"
	"Product-Service/dto"
	client "Product-Service/handler"
	"Product-Service/interfaces"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type productService struct {
	ctx        context.Context
	repo       interfaces.ProductInterface
	userClient *client.UserClient
	rdb        interfaces.ProductRedis
}

func NewProductService(repo interfaces.ProductInterface, uc *client.UserClient, redis interfaces.ProductRedis) interfaces.ProductService {
	return &productService{ctx: context.Background(), repo: repo, userClient: uc, rdb: redis}
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

	isSuccess := s.rdb.SetProductToCache("products", created)
	if !isSuccess {
		return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Create product successfully but Failed to cache product", created)
	}

	return *helpers.StandardApiResponse(true, http.StatusOK, "Create Product Successfully", created)
}

func (s *productService) GetByIDProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[dto.GetProductResponse] {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusBadRequest, "Invalid ID parameter", dto.GetProductResponse{})
	}
	var product dto.GetProductResponse
	product = s.rdb.GetProductByIdFromCache("products", idStr)
	log.Println("Product retrieve from cache")
	if (product == dto.GetProductResponse{}) {
		product = s.repo.GetProductByID(id)
		log.Println("Product retrieve from database")
	}

	return *helpers.StandardApiResponse(true, http.StatusOK, "Product retrieved successfully", product)
}

func (s *productService) UpdateProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[dto.GetProductResponse] {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	var updateProduct dto.UpdateProductDTO
	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusBadRequest, "Invalid ID parameter", dto.GetProductResponse{})
	}

	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusBadRequest, "Invalid request body", dto.GetProductResponse{})
	}

	updated := repo.UpdateProduct(id, updateProduct)
	if (updated == dto.GetProductResponse{}) {
		return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Failed to update product", dto.GetProductResponse{})
	}

	isExistKey := s.rdb.IsExistKeyInCache("products")
	if isExistKey {
		isSuccess := s.rdb.SetProductToCache("products", updated)
		if !isSuccess {
			return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Update product successfully but Failed to update cache product", updated)
		}
	}

	return *helpers.StandardApiResponse(true, http.StatusOK, "Product update successfully", updated)
}

func (s *productService) DeleteProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[dto.GetProductResponse] {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return *helpers.StandardApiResponse(false, http.StatusBadRequest, "Invalid ID parameter", dto.GetProductResponse{})
	}

	deletedProduct := repo.GetProductByID(id)
	if (deletedProduct == dto.GetProductResponse{}) {
		return *helpers.StandardApiResponse(false, http.StatusNotFound, "Product not found", dto.GetProductResponse{})
	}

	isDeleted := repo.DeleteProduct(id)
	if !isDeleted {
		return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Failed to delete product", dto.GetProductResponse{})
	}

	isExistKey := s.rdb.IsExistKeyInCache("products")
	if isExistKey {
		isSuccess := s.rdb.DeleteProductFromCache("products", strconv.Itoa(deletedProduct.Id))
		if !isSuccess {
			return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Delete product successfully but Failed to delete from cache", deletedProduct)
		}
	}

	return *helpers.StandardApiResponse(false, http.StatusOK, "Delete Product Successfully", deletedProduct)
}

func (s *productService) GetProductService(r *http.Request, repo interfaces.ProductInterface) helpers.ApiResponse[[]dto.GetProductResponse] {
	var products []dto.GetProductResponse
	productsFromCache, err := s.rdb.GetProductFromCache("products")
	if productsFromCache == nil || err != nil {
		productsFromDb, err := repo.GetProducts()
		if err != nil {
			return *helpers.StandardApiResponse(false, http.StatusInternalServerError, "Failed to get products", []dto.GetProductResponse{})
		}
		for _, p := range productsFromDb {
			isSuccess := s.rdb.SetProductToCache("products", p)
			if !isSuccess {
				log.Println("Failed to cache product with Id:", p.Id)
			}
		}
		products = append(products, productsFromDb...)
	} else {
		products = append(products, productsFromCache...)
	}
	return *helpers.StandardApiResponse(true, http.StatusOK, "Products retrieved successfully", products)
}
