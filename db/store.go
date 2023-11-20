package db

import (
	"context"
	"database/sql"
	"fmt"
)

type TRX interface {
	CreateAccount(ctx context.Context, a AccountParams) (Account, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	ListAccount(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParam) (Account, error)
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
}

type STRX struct {
	db   *Db
	conn *sql.DB
}

func NewTRX(conn *sql.DB) *STRX {
	return &STRX{
		db:   NewDb(conn),
		conn: conn,
	}
}

func (strx *STRX) execTx(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := strx.conn.Begin()
	if err != nil {
		return err
	}

	err = fn(tx)

	if err != nil {
		if errRbllbck := tx.Rollback(); errRbllbck != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, errRbllbck)
		}
		return err
	}

	return tx.Commit()
}
