package db

import (
	"context"
	"database/sql"
)

type TRX interface {
	CreateAccount(ctx context.Context, a AccountParams) (Account, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	ListAccount(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParam) (Account, error)
}

type STRX struct {
	db *Db
}

func NewTRX(conn *sql.DB) *STRX {
	return &STRX{
		db: NewDb(conn),
	}
}
