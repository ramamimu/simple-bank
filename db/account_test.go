package db

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	globalOwner = "bakyun"
)

func TestAccount(t *testing.T) {
	t.Run("test create account", func(t0 *testing.T) {
		act := AccountTestComunal{t}
		act.CreateAccount()
		act.GetAccount()
		act.UpdateAccount()
	})

	t.Run("test list account", func(t1 *testing.T) {
		act := AccountTestComunal{t}
		act.ListAccount()
	})
}

type AccountTestComunal struct {
	t *testing.T
}

func (atc *AccountTestComunal) CreateAccount() {
	a := AccountParams{
		Owner:     globalOwner,
		Balance:   0,
		Currency:  "1209",
		CreatedAt: time.Now(),
	}
	for i := 0; i < 3; i++ {
		a.Balance = rand.New(rand.NewSource(int64(i))).Int63()
		account, err := trx.CreateAccount(context.Background(), a)
		assert.NoError(atc.t, err)
		assert.NotNil(atc.t, account)
		assert.NotNil(atc.t, account.ID)
		assert.Equal(atc.t, account.Balance, a.Balance)
		assert.Equal(atc.t, account.Owner, a.Owner)
		assert.Equal(atc.t, account.Currency, a.Currency)
		assert.NotNil(atc.t, account.CreatedAt)
	}
}

func (atc *AccountTestComunal) GetAccount() {
	validID := []int64{1, 2, 3}

	for _, i := range validID {
		a, err := trx.GetAccount(context.Background(), i)
		assert.NoError(atc.t, err)
		assert.NotNil(atc.t, a)
	}

	ErrorID := []int64{1342432, 232432, 3423443}
	for _, i := range ErrorID {
		_, err := trx.GetAccount(context.Background(), i)
		assert.Error(atc.t, err)
	}
}

func (atc *AccountTestComunal) ListAccount() {
	p := ListAccountsParams{
		Owner:  globalOwner,
		Limit:  10,
		Offset: 2,
	}

	l, err := trx.ListAccount(context.Background(), p)

	assert.NoError(atc.t, err)
	assert.True(atc.t, len(l) >= defaultAccount)
}

func (atc *AccountTestComunal) UpdateAccount() {
	p := []UpdateAccountParam{
		{
			ID:      1,
			Balance: 11000,
		},
		{
			ID:      2,
			Balance: 12000,
		},
		{
			ID:      3,
			Balance: 13000,
		},
	}

	for _, i := range p {
		a, err := trx.UpdateAccount(context.Background(), i)
		assert.NoError(atc.t, err)
		assert.Equal(atc.t, i.Balance, a.Balance)
		assert.Equal(atc.t, i.ID, a.ID)
	}
}
