package route

import (
	"Auth-Service/handler"
	"Auth-Service/interfaces"
	"github.com/gorilla/mux"
)

func Route(repo interfaces.AuthRepository) *mux.Router {
	router := mux.NewRouter()
	h := handler.NewHandler(repo)

	router.HandleFunc("/registration", h.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", h.LoginHander).Methods("POST")

	return router
}
