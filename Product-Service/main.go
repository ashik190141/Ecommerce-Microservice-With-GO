package main

import (
	"Product-Service/config"
	"log"
	"net"
	"net/http"
	"Product-Service/app/db"
	grpc_client "Product-Service/handler"
	repo "Product-Service/repo"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadEnvData()
	address := ":" + cfg.ProductServicePort

	databse:= db.ConnectPostgres()
	repo.NewProductRepository(databse)
	userClient := grpc_client.NewUserClient(cfg.GrpcUserServiceUrl)

	grpcServer:= grpc.NewServer()
	grpc_client.NewProductGRPCServer(userClient)

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
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}