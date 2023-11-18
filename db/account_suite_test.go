package db

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type AccountTest struct {
	suite.Suite
	db *sql.DB
}

func TestAccountTest(t *testing.T) {
	suite.Run(t, &AccountTest{})
}

func (at *AccountTest) SetupSuite() {
	at.db = conn
}

func (at *AccountTest) TestAccount() {
	a := AccountParams{
		Owner:     globalOwner,
		Balance:   0,
		Currency:  "1209",
		CreatedAt: time.Now(),
	}
	for i := 0; i < 3; i++ {
		a.Balance = rand.New(rand.NewSource(int64(i))).Int63()
		account, err := trx.CreateAccount(context.Background(), a)
		at.NoError(err)
		at.NotNil(account)
		at.NotNil(account.ID)
		at.Equal(account.Balance, a.Balance)
		at.Equal(account.Owner, a.Owner)
		at.Equal(account.Currency, a.Currency)
		at.NotNil(account.CreatedAt)
	}
}
