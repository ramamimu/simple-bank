package db

import (
	"context"
	"database/sql"
)

type TRX interface {
	CreateAccount(ctx context.Context, a Account) (Account, error)
}

type STRX struct {
	db *Db
	TRX
}

func NewTRX(conn *sql.DB) *STRX {
	return &STRX{
		db: NewDb(conn),
	}
}
