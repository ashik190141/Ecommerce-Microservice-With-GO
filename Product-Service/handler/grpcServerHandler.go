package handler

import (
    pb "Product-Service/grpc/user/grpc"
)

type ProductGRPCServer struct {
    pb.UnimplementedProductServiceServer
	UserClient *UserClient
}

func NewProductGRPCServer(userClient *UserClient) *ProductGRPCServer {
    return &ProductGRPCServer{
		UserClient: userClient,
	}
}
