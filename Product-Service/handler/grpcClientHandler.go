package handler

import (
	userPb "Product-Service/grpc/user/grpc"
	"context"
	"log"

	"google.golang.org/grpc"
)

type UserClient struct {
	client userPb.ProductServiceClient
}

func NewUserClient(authServiceURL string) *UserClient {
	conn, err := grpc.Dial(authServiceURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to UserService: %v", err)
	}

	return &UserClient{
		client: userPb.NewProductServiceClient(conn),
	}
}

func (uc *UserClient) IsUserExist(ctx context.Context, email string) (*userPb.IsUserExistResponse, error) {
	resp, err := uc.client.IsUserExist(ctx, &userPb.IsUserExistRequest{Email: email})
	if err != nil {
		log.Printf("gRPC IsUserExist() failed: %v", err)
		return nil, err
	}

	return resp, nil
}