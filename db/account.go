package db

import (
	"context"
	"time"
)

type AccountParams struct {
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

const createAccount = `
INSERT INTO accounts (owner, balance, currency, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *`

func (s *STRX) CreateAccount(ctx context.Context, a AccountParams) (Account, error) {
	var acc Account
	err := s.db.QueryRowContext(ctx, createAccount, a.Owner, a.Balance, a.Currency, a.CreatedAt).Scan(&acc.ID, &acc.Owner, &acc.Balance, &acc.Currency, &acc.CreatedAt)

	if err != nil {
		return acc, err
	}

	return acc, nil
}

const getAccount = `
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE id = $1 LIMIT 1
`

func (s *STRX) GetAccount(ctx context.Context, id int64) (Account, error) {
	var acc Account
	err := s.db.QueryRowContext(ctx, getAccount, id).Scan(&acc.ID, &acc.Owner, &acc.Balance, &acc.Currency, &acc.CreatedAt)

	return acc, err
}

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

const listAccounts = `
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

func (s *STRX) ListAccount(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := s.db.QueryContext(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Account{}

	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
