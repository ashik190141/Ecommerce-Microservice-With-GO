package main

import (
	"Product-Service/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadEnvData()
	address := ":" + cfg.ProductServicePort
	fmt.Println("Server running on port", cfg.ProductServicePort)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}