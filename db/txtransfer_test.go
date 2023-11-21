package db

import (
	"context"
	"fmt"
	"time"
)

func (ac *AccountTest) TestTransferTx() {

	n := 20
	// synchronous transfer
	for i := 0; i < n; i++ {
		id1 := int64((i + 1) % defaultAccount)
		id2 := int64((i + 3) % defaultAccount)
		if id1 == 0 {
			id1++
		} else if id2 == 0 {
			id2++
		}

		preAccFrom, errPreAccFrom := trx.GetAccount(context.Background(), id1)
		preAccTo, errPreAccTo := trx.GetAccount(context.Background(), id2)
		ac.NoError(errPreAccFrom)
		ac.NoError(errPreAccTo)

		amount := int64(((i + 1) * 100) % 38)

		tx, err := trx.TransferTx(context.Background(), TransferTxParams{
			FromAccountID: id1,
			ToAccountID:   id2,
			Amount:        amount,
		})

		accFrom, errAccFrom := trx.GetAccount(context.Background(), id1)
		accTo, errAccTo := trx.GetAccount(context.Background(), id2)

		ac.NoError(err)
		ac.NoError(errAccFrom)
		ac.NoError(errAccTo)

		ac.Equal(tx.FromAccount.Balance+amount, preAccFrom.Balance)
		ac.Equal(tx.FromAccount.ID, preAccFrom.ID)
		ac.Equal(tx.FromAccount.Owner, preAccFrom.Owner)
		ac.Equal(tx.FromAccount.Currency, preAccFrom.Currency)
		ac.Equal(tx.FromAccount.CreatedAt, preAccFrom.CreatedAt)

		ac.Equal(tx.ToAccount.Balance-amount, preAccTo.Balance)
		ac.Equal(tx.ToAccount.ID, preAccTo.ID)
		ac.Equal(tx.ToAccount.Owner, preAccTo.Owner)
		ac.Equal(tx.ToAccount.Currency, preAccTo.Currency)
		ac.Equal(tx.ToAccount.CreatedAt, preAccTo.CreatedAt)

		ac.Equal(preAccFrom.Balance-amount, accFrom.Balance)
		ac.Equal(preAccFrom.ID, accFrom.ID)
		ac.Equal(preAccFrom.Owner, accFrom.Owner)
		ac.Equal(preAccFrom.Currency, accFrom.Currency)
		ac.Equal(preAccFrom.CreatedAt, accFrom.CreatedAt)

		ac.Equal(preAccTo.Balance+amount, accTo.Balance)
		ac.Equal(preAccTo.ID, accTo.ID)
		ac.Equal(preAccTo.Owner, accTo.Owner)
		ac.Equal(preAccTo.Currency, accTo.Currency)
		ac.Equal(preAccTo.CreatedAt, accTo.CreatedAt)

	}
}

func (ac *AccountTest) TestAsyncTransferTx() {

	n := 20

	// test deadlock potential
	errTx := make(chan error)

	// asynchronous transfer
	for i := 0; i < n; i++ {
		j := i
		go func() {
			// process <- j
			id1 := int64((j + 1) % defaultAccount)
			id2 := int64((j + 3) % defaultAccount)
			if id1 == 0 {
				id1++
			} else if id2 == 0 {
				id2++
			}

			amount := int64(((j + 1) * 100) % 38)

			time.Sleep((time.Duration(amount)%10 + 1) * time.Second)
			_, err := trx.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: id1,
				ToAccountID:   id2,
				Amount:        amount,
			})
			errTx <- err

			ac.NoError(err)
		}()
	}
	for i := 0; i < n; i++ {
		err := <-errTx
		if err != nil {
			fmt.Println(err)
		}
		ac.NoError(err)
	}
}

func (ac *AccountTest) TestOutOfRange() {
	account, err := trx.CreateAccount(context.Background(), AccountParams{
		Owner:    "kasuari",
		Balance:  99,
		Currency: "IDR",
	})
	ac.NoError(err)

	_, errAccBalance := trx.AddAccountBalance(context.Background(), AddAccountBalanceParams{
		ID:     account.ID,
		Amount: -100,
	})
	ac.Error(errAccBalance)

	preAddBalance, errPreAddBalance := trx.GetAccount(context.Background(), account.ID)
	amount := 100
	accBalance2, errAccBalance2 := trx.AddAccountBalance(context.Background(), AddAccountBalanceParams{
		ID:     account.ID,
		Amount: int64(amount),
	})
	ac.NoError(errPreAddBalance)
	ac.NoError(errAccBalance2)
	ac.Equal(preAddBalance.ID, account.ID)
	ac.Equal(preAddBalance.Currency, accBalance2.Currency)
	ac.Equal(accBalance2.Balance, preAddBalance.Balance+int64(amount))
}
