package route

import (
	"Auth-Service/handler"
	"Auth-Service/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

func Route(repo interfaces.AuthRepository) *mux.Router {
	router := mux.NewRouter()
	h := handler.NewHandler(repo)

	router.HandleFunc("/registration", h.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", h.LoginHander).Methods("POST")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Server is running 🚀"))
	}).Methods("GET")

	return router
}
