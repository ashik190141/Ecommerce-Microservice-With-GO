package main

import (
	"Auth-Service/app/db"
	"Auth-Service/app/route"
	"Auth-Service/config"
	pb "Auth-Service/grpc/grpc"
	"Auth-Service/handler"
	"Auth-Service/repo"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadEnvData()
	address := ":" + cfg.AuthServicePort
	fmt.Println("Server running on port", cfg.AuthServicePort)

	database := db.GetDBConnection()
	userRepo := repo.NewUserRepository(database)

	grpcServer := grpc.NewServer()
	authService := handler.NewProductGRPCServer(userRepo)
	pb.RegisterProductServiceServer(grpcServer, authService)

	router := route.Route(userRepo)

	lis, err := net.Listen("tcp", ":"+cfg.GRPCAuthServicePort)
	if err != nil {
		log.Fatalf("❌ Failed to listen on gRPC port %s: %v", cfg.GRPCAuthServicePort, err)
	}

	log.Printf("🚀 gRPC Auth Service running on port %s", cfg.GRPCAuthServicePort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve gRPC: %v", err)
	}

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
