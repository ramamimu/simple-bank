package db

import (
	"context"
)

func (ctt *AccountTest) TestCreateTransfer() {
	ctp := []CreateTransferParams{
		{
			FromAccountID: 1,
			ToAccountID:   2,
			Amount:        10,
		},
		{
			FromAccountID: 3,
			ToAccountID:   4,
			Amount:        20,
		},
	}

	for _, i := range ctp {
		tx, err := trx.CreateTransfer(context.Background(), i)
		ctt.NoError(err)
		ctt.Equal(tx.FromAccountID, i.FromAccountID)
		ctt.Equal(tx.ToAccountID, i.ToAccountID)
		ctt.Equal(tx.Amount, i.Amount)
	}
}
