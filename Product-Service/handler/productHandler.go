package handler

import (
	"Product-Service/interfaces"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	service interfaces.ProductService
	repo   interfaces.ProductInterface
}

func NewProductHandler(service interfaces.ProductService, repo interfaces.ProductInterface) interfaces.ProductHandler {
	return &ProductHandler{service: service, repo: repo}
}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	response := h.service.CreateProductService(r, h.repo)
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	response := h.service.GetProductService(r, h.repo)
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	response := h.service.GetByIDProductService(r, h.repo)
	json.NewEncoder(w).Encode(response)
}