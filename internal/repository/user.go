package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User interface {
	CreateUser(ctx context.Context) (int64, error)
}

type user struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) User {
	return &user{
		db: db,
	}
}

func (u *user) CreateUser(ctx context.Context) (int64, error) {
	panic("implement me")
}
