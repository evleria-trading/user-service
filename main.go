package main

import (
	"context"
	"fmt"
	"github.com/evleria-trading/user-service/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func main() {
	cfg := new(config.Config)

	db := getPostgres(cfg)
	defer db.Close()

	//userRepository := repository.NewUserRepository(db)

}

func getPostgres(cfg *config.Config) *pgxpool.Pool {
	dbURL := getPostgresConnectionString(cfg)

	db, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getPostgresConnectionString(cfg *config.Config) string {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDb,
	)
	if cfg.PostgresSSLDisable {
		conn += "?sslmode=disable"
	}
	return conn
}
