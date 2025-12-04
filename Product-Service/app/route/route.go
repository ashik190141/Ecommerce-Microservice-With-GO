package route

import (
	handler "Product-Service/handler"
	"Product-Service/interfaces"
	"net/http"
	"github.com/gorilla/mux"
)

func Route(repo interfaces.ProductInterface, service interfaces.ProductService) *mux.Router {
	router := mux.NewRouter()
	h := handler.NewProductHandler(service, repo)

	router.HandleFunc("/create-product", h.CreateProductHandler).Methods("POST")
	router.HandleFunc("/get-product", h.GetProductsHandler).Methods("GET")
	router.HandleFunc("/get-product/{id}", h.GetProductByIDHandler).Methods("GET")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Server is running"))
	}).Methods("GET")

	return router
}
