package db

import "context"

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (s *STRX) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	const createTransfer = `-- name: CreateTransfer :one
	INSERT INTO transfers (
	  from_account_id,
	  to_account_id,
	  amount
	) VALUES (
	  $1, $2, $3
	) RETURNING id, from_account_id, to_account_id, amount, created_at
	`

	row := s.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
