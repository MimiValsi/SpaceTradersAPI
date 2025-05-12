package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn *pgx.Conn
}

func InitDB(url string) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	db := DB{
		conn: conn,
	}
	return &db, nil
}
