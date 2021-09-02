package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evleria-trading/user-service/internal/config"
	"github.com/evleria-trading/user-service/internal/handler"
	"github.com/evleria-trading/user-service/internal/repository"
	"github.com/evleria-trading/user-service/internal/service"
	"github.com/evleria-trading/user-service/protocol/pb"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	cfg := new(config.Config)
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db := getPostgres(cfg)
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	startGrpcServer(handler.NewUserService(userService), ":6000")
}

func startGrpcServer(userService pb.UserServiceServer, port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)
	reflection.Register(s)

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
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
