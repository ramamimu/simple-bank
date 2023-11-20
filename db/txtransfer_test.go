package db

import (
	"context"
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

// func (ac *AccountTest) TestAsyncTransferTx() {

// 	n := 2

// 	errTx := make(chan error)
// 	trfTx := make(chan TransferTxResult)
// 	accFromTx := make(chan Account)
// 	accToTx := make(chan Account)
// 	preAccFromTx := make(chan Account)
// 	preAccToTx := make(chan Account)
// 	amountTx := make(chan int64)
// 	process := make(chan int)
// 	// asynchronous transfer
// 	for i := 0; i < n; i++ {
// 		j := i
// 		go func() {
// 			process <- j
// 			id1 := int64((j + 1) % defaultAccount)
// 			id2 := int64((j + 3) % defaultAccount)
// 			if id1 == 0 {
// 				id1++
// 			} else if id2 == 0 {
// 				id2++
// 			}

// 			fmt.Printf("from %d, to %d\n\n", id1, id2)

// 			preAccFrom, errPreAccFrom := trx.GetAccount(context.Background(), id1)
// 			preAccTo, errPreAccTo := trx.GetAccount(context.Background(), id2)
// 			ac.NoError(errPreAccFrom)
// 			ac.NoError(errPreAccTo)

// 			amount := int64(((j + 1) * 100) % 38)

// 			tx, err := trx.TransferTx(context.Background(), TransferTxParams{
// 				FromAccountID: id1,
// 				ToAccountID:   id2,
// 				Amount:        amount,
// 			})

// 			accFrom, errAccFrom := trx.GetAccount(context.Background(), id1)
// 			accTo, errAccTo := trx.GetAccount(context.Background(), id2)

// 			ac.NoError(errAccFrom)
// 			ac.NoError(errAccTo)

// 			errTx <- err
// 			trfTx <- tx
// 			accFromTx <- accFrom
// 			accToTx <- accTo
// 			preAccFromTx <- preAccFrom
// 			preAccToTx <- preAccTo
// 			amountTx <- amount
// 		}()
// 	}

// 	for i := 0; i < n; i++ {
// 		fmt.Println("Process ", <-process)
// 		amount := <-amountTx
// 		err := <-errTx
// 		trf := <-trfTx

// 		ac.NoError(err)

// 		preAccFrom := <-preAccFromTx
// 		ac.Equal(trf.FromAccount.Balance+amount, preAccFrom.Balance)
// 		ac.Equal(trf.FromAccount.ID, preAccFrom.ID)
// 		ac.Equal(trf.FromAccount.Owner, preAccFrom.Owner)
// 		ac.Equal(trf.FromAccount.Currency, preAccFrom.Currency)
// 		ac.Equal(trf.FromAccount.CreatedAt, preAccFrom.CreatedAt)

// 		preAccTo := <-preAccToTx
// 		ac.Equal(trf.ToAccount.Balance-amount, preAccTo.Balance)
// 		ac.Equal(trf.ToAccount.ID, preAccTo.ID)
// 		ac.Equal(trf.ToAccount.Owner, preAccTo.Owner)
// 		ac.Equal(trf.ToAccount.Currency, preAccTo.Currency)
// 		ac.Equal(trf.ToAccount.CreatedAt, preAccTo.CreatedAt)

// 		accFrom := <-accFromTx
// 		ac.Equal(trf.FromAccount.Balance-amount, accFrom.Balance)
// 		ac.Equal(trf.FromAccount.ID, accFrom.ID)
// 		ac.Equal(trf.FromAccount.Owner, accFrom.Owner)
// 		ac.Equal(trf.FromAccount.Currency, accFrom.Currency)
// 		ac.Equal(trf.FromAccount.CreatedAt, accFrom.CreatedAt)

// 		accTo := <-accToTx
// 		ac.Equal(trf.FromAccount.Balance+amount, accTo.Balance)
// 		ac.Equal(trf.FromAccount.ID, accTo.ID)
// 		ac.Equal(trf.FromAccount.Owner, accTo.Owner)
// 		ac.Equal(trf.FromAccount.Currency, accTo.Currency)
// 		ac.Equal(trf.FromAccount.CreatedAt, accTo.CreatedAt)

// 		fmt.Println("Process Finished")
// 	}
// }
