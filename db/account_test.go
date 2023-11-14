package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	a := Account{
		Owner:     "bakyun",
		Balance:   1000,
		Currency:  "1209",
		CreatedAt: time.Now(),
	}
	account, err := trx.CreateAccount(context.Background(), a)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.NotNil(t, account.ID)
	assert.Equal(t, account.Balance, a.Balance)
	assert.Equal(t, account.Owner, a.Owner)
	assert.Equal(t, account.Currency, a.Currency)
	assert.NotNil(t, account.CreatedAt)
}
