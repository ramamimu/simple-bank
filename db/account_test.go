package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	a := AccountParams{
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

func TestGetAccount(t *testing.T) {
	validID := []int64{1, 2, 3}

	for _, i := range validID {
		a, err := trx.GetAccount(context.Background(), i)
		assert.NoError(t, err)
		assert.NotNil(t, a)
	}

	ErrorID := []int64{1342432, 232432, 3423443}
	for _, i := range ErrorID {
		_, err := trx.GetAccount(context.Background(), i)
		assert.Error(t, err)
	}
}
