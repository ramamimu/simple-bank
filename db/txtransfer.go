package db

import (
	"context"
	"database/sql"
	"fmt"
)

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (s *STRX) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var tfr TransferTxResult
	err := s.execTx(ctx, func(tx *sql.Tx) error {
		// add transfer
		// add entries to account
		// add entries from account
		// update to account
		// update from account
		fmt.Println("Hello World")
		return nil
	})
	return tfr, err
}
