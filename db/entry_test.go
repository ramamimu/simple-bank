package db

import "context"

func (at *AccountTest) TestCreateEntry() {
	cep := []CreateEntryParams{
		{
			AccountID: 1,
			Amount:    100,
		},
		{
			AccountID: 2,
			Amount:    -100,
		},
	}

	for _, i := range cep {
		ce, err := trx.CreateEntry(context.Background(), i)
		at.Equal(ce.AccountID, i.AccountID)
		at.Equal(ce.Amount, i.Amount)
		at.NoError(err)
	}
}
