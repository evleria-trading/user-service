package service

import (
	"context"
	"errors"
	"github.com/evleria-trading/user-service/internal/repository"
)

var (
	ErrBalanceIsNegative = errors.New("balance is negative")
	ErrUserNotFound      = errors.New("user not found")
)

type User interface {
	CreateUser(ctx context.Context) (int64, error)
	SetBalance(ctx context.Context, balance float64, id int64) error
	GetBalanceByID(ctx context.Context, id int64) (float64, error)
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

func (u *user) SetBalance(ctx context.Context, balance float64, id int64) error {
	if balance < 0 {
		return ErrBalanceIsNegative
	}
	err := u.userRepository.UpdateBalance(ctx, balance, id)
	if err == repository.ErrUserNotFound {
		return ErrUserNotFound
	}
	return err
}

func (u *user) GetBalanceByID(ctx context.Context, id int64) (float64, error) {
	bal, err := u.userRepository.GetBalance(ctx, id)
	if err != nil {
		return 0, err
	}
	return bal, nil
}
