package db

import "context"

func (s *STRX) CreateAccount(ctx context.Context, a Account) (Account, error) {
	sqlStatement := `
		INSERT INTO accounts (owner, balance, currency, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING *`

	var acc Account
	err := s.db.QueryRowContext(context.Background(), sqlStatement, a.Owner, a.Balance, a.Currency, a.CreatedAt).Scan(&acc.ID, &acc.Owner, &acc.Balance, &acc.Currency, &acc.CreatedAt)

	if err != nil {
		return acc, err
	}

	return acc, nil
}
