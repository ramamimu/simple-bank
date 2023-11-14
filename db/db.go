package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "postgres"
	dbname   string = "simplebank"
)

func LoadConfigDb() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

type Queries interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type Db struct {
	Queries
}

func NewDb(conn *sql.DB) *Db {
	return &Db{
		conn,
	}
}
