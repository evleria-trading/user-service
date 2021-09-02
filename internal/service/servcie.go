package service

import (
	"context"
	"github.com/evleria-trading/user-service/protocol/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User interface {
}

type user struct {
}

func NewUserService() User {
	return &user{}
}

func (u *user) CreateUser(ctx context.Context) (*pb.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (u *user) SetBalance(ctx context.Context, request *pb.SetBalanceRequest) error {
	return status.Errorf(codes.Unimplemented, "method SetBalance not implemented")
}
