package handler

import (
	"Auth-Service/interfaces"
	"Auth-Service/service/auth"
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo interfaces.AuthRepository
}

func NewHandler(repo interfaces.AuthRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	response := auth.CreateUserService(r, h.repo)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) LoginHander(w http.ResponseWriter, r *http.Request) {
	response := auth.AuthLoginService(r, h.repo)
	json.NewEncoder(w).Encode(response)
}
