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

func (u *UserService) CreateUser(ctx context.Context, _ *empty.Empty) (*pb.CreateUserResponse, error) {
	id, err := u.service.CreateUser(ctx)
	if err != nil {
		return nil, status.Error(getStatusCode(err), err.Error())
	}

	return &pb.CreateUserResponse{
		UserId: id,
	}, nil
}

func (u *UserService) SetBalance(ctx context.Context, request *pb.SetBalanceRequest) (*empty.Empty, error) {
	err := u.service.SetBalance(ctx, request.Balance, request.UserId)
	if err != nil {
		return nil, status.Error(getStatusCode(err), err.Error())
	}

	return &empty.Empty{}, nil
}

func (u *UserService) GetBalance(ctx context.Context, request *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	bal, err := u.service.GetBalanceByID(ctx, request.UserId)
	if err != nil {
		return nil, status.Error(getStatusCode(err), err.Error())
	}
	return &pb.GetBalanceResponse{
		Balance: bal,
	}, nil
}

func getStatusCode(err error) codes.Code {
	switch err {
	case service.ErrUserNotFound:
		return codes.NotFound
	case service.ErrBalanceIsNegative:
		return codes.InvalidArgument
	default:
		return codes.Internal
	}
}
