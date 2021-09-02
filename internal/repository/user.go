package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrUserNotFound = errors.New("user not found")

type User interface {
	CreateUser(ctx context.Context) (int64, error)
	UpdateBalance(ctx context.Context, balance float64, id int64) error
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
	var id int64
	err := u.db.QueryRow(ctx, `INSERT INTO users DEFAULT VALUES RETURNING user_id;`).Scan(&id)
	return id, err
}

func (u *user) UpdateBalance(ctx context.Context, balance float64, id int64) error {
	res, err := u.db.Exec(ctx, `UPDATE users SET balance=$1 WHERE user_id=$2;`, balance, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrUserNotFound
	}
	return nil
}
