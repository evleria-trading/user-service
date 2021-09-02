package handler

import "github.com/evleria-trading/user-service/protocol/pb"

type UserService struct {
	pb.UserServiceServer
}

func NewUserService() pb.UserServiceServer {
	return &UserService{}
}
