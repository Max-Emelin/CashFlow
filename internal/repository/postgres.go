package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	DBUser   string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		logrus.Fatalf("Unable to create connection pool: %s\n", err.Error())
		os.Exit(1)
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
