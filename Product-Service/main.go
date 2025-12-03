package main

import (
	"Product-Service/app/db"
	"Product-Service/app/route"
	"Product-Service/config"
	grpc_client "Product-Service/handler"
	repo "Product-Service/repo"
	"log"
	"net"
	"net/http"
	service "Product-Service/service"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadEnvData()
	address := ":" + cfg.ProductServicePort

	database:= db.ConnectPostgres()
	productRepo := repo.NewProductRepository(database)
	userClient := grpc_client.NewUserClient(cfg.GrpcUserServiceUrl)
	productSrv := service.NewProductService(productRepo, userClient)

	grpcServer:= grpc.NewServer()
	grpc_client.NewProductGRPCServer(userClient)

	router := route.Route(productRepo, productSrv)

	lis, err := net.Listen("tcp", ":"+cfg.GRPCProductServicePort)
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port %s: %v", cfg.GRPCProductServicePort, err)
	}

	log.Printf("gRPC Product Service running on port %s", cfg.GRPCProductServicePort)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
	}()

	log.Printf("HTTP Product Service running on port %s", cfg.ProductServicePort)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}