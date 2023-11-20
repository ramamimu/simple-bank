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
	var tfr TransferTxResult
	err := s.execTx(ctx, func(tx *sql.Tx) error {
		// add transfer
		_, err := s.CreateTransfer(context.Background(), CreateTransferParams(arg))

		if err != nil {
			return err
		}

		// add entries from account
		entryFrom, errEntryFrom := s.CreateEntry(context.Background(), CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if errEntryFrom != nil {
			return err
		}

		// add entries to account
		entryTo, errEntryTo := s.CreateEntry(context.Background(), CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    -arg.Amount,
		})

		if errEntryTo != nil {
			return err
		}

		// update from account
		_, errUpdateFrom := s.UpdateAccount(context.Background(), UpdateAccountParam{
			ID:      arg.FromAccountID,
			Balance: entryFrom.Amount - arg.Amount,
		})
		if errUpdateFrom != nil {
			return err
		}

		// update to account
		_, errUpdateTo := s.UpdateAccount(context.Background(), UpdateAccountParam{
			ID:      arg.ToAccountID,
			Balance: entryTo.Amount + arg.Amount,
		})
		if errUpdateTo != nil {
			return err
		}

		return nil
	})
	return tfr, err
}
