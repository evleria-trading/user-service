package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrUserNotFound = errors.New("user not found")

type User interface {
	CreateUser(ctx context.Context) (int64, error)
	UpdateBalance(ctx context.Context, balance float64, id int64) error
	GetBalance(ctx context.Context, id int64) (float64, error)
	AddBalance(ctx context.Context, id int64, diff float64) (float64, error)
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

func (u *user) GetBalance(ctx context.Context, id int64) (float64, error) {
	var bal float64
	err := u.db.QueryRow(ctx, `SELECT balance FROM users WHERE user_id=$1;`, id).Scan(&bal)

	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, ErrUserNotFound
		}
		return 0, err
	}
	return bal, nil
}

func (u *user) AddBalance(ctx context.Context, id int64, diff float64) (float64, error) {
	var bal float64
	err := u.db.QueryRow(ctx, `UPDATE users SET balance=balance + $1 WHERE user_id=$2 RETURNING balance;`, diff, id).Scan(&bal)

	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, ErrUserNotFound
		}
		return 0, err
	}
	return bal, nil
}
