package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	connect, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s username=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = connect.Ping()
	if err != nil {
		return nil, err
	}

	return connect, nil
}
