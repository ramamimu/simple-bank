package db

import (
	"context"
	"fmt"
	"time"
)

type AccountParams struct {
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *STRX) CreateAccount(ctx context.Context, a AccountParams) (Account, error) {
	const createAccount = `
	INSERT INTO accounts (owner, balance, currency, created_at)
	VALUES ($1, $2, $3, $4)
	RETURNING *`

	var acc Account
	err := s.db.QueryRowContext(ctx, createAccount, a.Owner, a.Balance, a.Currency, a.CreatedAt).Scan(&acc.ID, &acc.Owner, &acc.Balance, &acc.Currency, &acc.CreatedAt)

	if err != nil {
		return acc, err
	}

	return acc, nil
}

func (s *STRX) GetAccount(ctx context.Context, id int64) (Account, error) {
	const getAccount = `
	SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE id = $1 LIMIT 1
	`
	var acc Account
	err := s.db.QueryRowContext(ctx, getAccount, id).Scan(&acc.ID, &acc.Owner, &acc.Balance, &acc.Currency, &acc.CreatedAt)

	return acc, err
}

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (s *STRX) ListAccount(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	const listAccounts = `
	SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE owner = $1
	ORDER BY id
	`

	rows, err := s.db.QueryContext(ctx, listAccounts, arg.Owner)
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

type UpdateAccountParam struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (s *STRX) UpdateAccount(ctx context.Context, arg UpdateAccountParam) (Account, error) {
	const updateAccount = `
	UPDATE accounts 
	SET balance = $2
	WHERE id = $1
	RETURNING id, owner, balance, currency, created_at 
	`
	row := s.db.QueryRowContext(ctx, updateAccount, arg.ID, arg.Balance)

	var a Account
	if err := row.Scan(
		&a.ID,
		&a.Owner,
		&a.Balance,
		&a.Currency,
		&a.CreatedAt,
	); err != nil {
		return a, err
	}
	return a, nil
}

const addAccountBalance = `-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + $1
WHERE id = $2
RETURNING id, owner, balance, currency, created_at
`

type AddAccountBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (s *STRX) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error) {
	row := s.conn.QueryRowContext(ctx, addAccountBalance, arg.Amount, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	if i.Balance < 0 {
		return i, fmt.Errorf("balance from id %d less than 0", i.ID)
	}
	return i, err
}
