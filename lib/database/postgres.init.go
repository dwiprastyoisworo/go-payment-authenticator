package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/config"
	_ "github.com/lib/pq"
	"time"
)

func PostgresInit(config *config.AppConfig, context context.Context) (*sql.Conn, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Database,
		config.Postgres.Ssl,
		config.Postgres.Schema,
	)

	// setup postgres connection
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	// setup postgres connection pool
	db.SetConnMaxIdleTime(time.Duration(config.Postgres.MaxIdleTime) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(config.Postgres.MaxLifeTime) * time.Minute)
	db.SetMaxOpenConns(config.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(config.Postgres.MaxIdleConns)

	conn, err := db.Conn(context)
	if err != nil {
		return nil, err
	}

	return conn, nil

}
