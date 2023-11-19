package db

import "context"

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (s *STRX) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	const createEntry = `-- name: CreateEntry :one
	INSERT INTO entries (
	  account_id,
	  amount
	) VALUES (
	  $1, $2
	) RETURNING id, account_id, amount, created_at
	`

	row := s.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
