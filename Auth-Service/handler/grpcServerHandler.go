package handler

import (
	pb "Auth-Service/grpc/grpc"
	"Auth-Service/interfaces"
	"context"
)

type AuthGRPCServer struct {
    pb.UnimplementedProductServiceServer
    repo interfaces.AuthRepository
}

func NewProductGRPCServer(repo interfaces.AuthRepository) *AuthGRPCServer {
    return &AuthGRPCServer{repo: repo}
}

func (s *AuthGRPCServer) IsUserExist(ctx context.Context, req *pb.IsUserExistRequest) (*pb.IsUserExistResponse, error) {
    exists, err := s.repo.GetUserByEmail(req.Email)
    if err != nil {
        return &pb.IsUserExistResponse{
            Success: false,
        }, nil
    }
    return &pb.IsUserExistResponse{
        Success: exists,
    }, nil
}