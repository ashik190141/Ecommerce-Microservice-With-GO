package main

import (
	"Auth-Service/app/db"
	"Auth-Service/app/route"
	"Auth-Service/config"
	"Auth-Service/repo"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadEnvData()
	address := ":" + cfg.AuthServicePort
	fmt.Println("Server running on port", cfg.AuthServicePort)

	database := db.GetDBConnection()
	userRepo := repo.NewUserRepository(database)

	router := route.Route(userRepo)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
