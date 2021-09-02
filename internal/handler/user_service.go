package handler

import (
	"context"
	"github.com/evleria-trading/user-service/internal/service"
	"github.com/evleria-trading/user-service/protocol/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	service service.User
}

func NewUserService(userService service.User) pb.UserServiceServer {
	return &UserService{
		service: userService,
	}
}

func (u *UserService) CreateUser(ctx context.Context, request *empty.Empty) (*pb.CreateUserResponse, error) {
	id, err := u.service.CreateUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unimplemented, err.Error())
	}
	return &pb.CreateUserResponse{
		UserId: id,
	}, nil
}

func (u *UserService) SetBalance(ctx context.Context, request *pb.SetBalanceRequest) (*empty.Empty, error) {
	err := u.service.SetBalance(ctx, request.Balance)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &empty.Empty{}, nil
}
