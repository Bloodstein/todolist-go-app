package repository

import (
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(connectionString string) (*sqlx.DB, error) {
	// give a dataSourceName argument as url-format string working
	connect, err := sqlx.Open("postgres", connectionString)

	// give a dataSourceName as string with params doesn't working
	// connect, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s name=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = connect.Ping()

	if err != nil {
		return nil, err
	}

	return connect, nil
}
