package main

import (
	"Product-Service/config"
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_client "Product-Service/handler"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadEnvData()
	address := ":" + cfg.ProductServicePort
	fmt.Println("Server running on port", cfg.ProductServicePort)

	userClient := grpc_client.NewUserClient(cfg.GrpcUserServiceUrl)

	grpcServer:= grpc.NewServer()
	grpc_client.NewProductGRPCServer(userClient)

	lis, err := net.Listen("tcp", ":"+cfg.GRPCProductServicePort)
	if err != nil {
		log.Fatalf("❌ Failed to listen on gRPC port %s: %v", cfg.GRPCProductServicePort, err)
	}

	log.Printf("🚀 gRPC Product Service running on port %s", cfg.GRPCProductServicePort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve gRPC: %v", err)
	}

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}