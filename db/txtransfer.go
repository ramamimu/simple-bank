package db

import (
	"context"
	"database/sql"
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
	var trf TransferTxResult
	// trf := make(chan TransferTxResult)
	err := s.execTx(context.Background(), func(tx *sql.Tx) error {
		// add transfer
		var err error
		trf.Transfer, err = s.CreateTransfer(context.Background(), CreateTransferParams(arg))

		if err != nil {
			return err
		}

		// add entries from account
		var errEntryFrom error
		trf.FromEntry, errEntryFrom = s.CreateEntry(context.Background(), CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if errEntryFrom != nil {
			return err
		}

		// add entries to account
		var errEntryTo error
		trf.ToEntry, errEntryTo = s.CreateEntry(context.Background(), CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if errEntryTo != nil {
			return err
		}

		// update from account
		var errUpdateFrom error
		trf.FromAccount, errUpdateFrom = s.AddAccountBalance(context.Background(), AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if errUpdateFrom != nil {
			return err
		}

		// update to account
		var errUpdateTo error
		trf.ToAccount, errUpdateTo = s.AddAccountBalance(context.Background(), AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})
		if errUpdateTo != nil {
			return err
		}

		return nil
	})

	return trf, err
}
