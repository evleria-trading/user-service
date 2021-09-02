package service

import (
	"context"
	"github.com/evleria-trading/user-service/internal/repository"
)

type User interface {
	CreateUser(ctx context.Context) (int64, error)
	SetBalance(ctx context.Context, balance float64) error
}

type user struct {
	userRepository repository.User
}

func NewUserService(userRepository repository.User) User {
	return &user{
		userRepository: userRepository,
	}
}

func (u *user) CreateUser(ctx context.Context) (int64, error) {
	id, err := u.userRepository.CreateUser(ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *user) SetBalance(ctx context.Context, balance float64) error {
	panic("implement me")
}
