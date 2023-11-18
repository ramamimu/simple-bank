package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
)

var trx TRX
var conn *sql.DB

func truncateTableAccount() {
	const truncateTableAccount = `
	TRUNCATE TABLE accounts CASCADE
	`
	_, err := conn.QueryContext(context.Background(), truncateTableAccount)
	if err != nil {
		log.Fatal("error when truncate table ", err)
	}

	const getAccount = `
	SELECT COUNT(*) FROM accounts
	`
	rowAccount := conn.QueryRowContext(context.Background(), getAccount)

	var total int
	errTotal := rowAccount.Scan(&total)
	if errTotal != nil || total != 0 {
		log.Fatal(func() string {
			if errTotal != nil {
				return "error when select to get total"
			} else if total != 0 {
				return "total is not zero"
			}

			return "error when select to get total and the total is not zero"
		}())
	}

	const resetSequence = `
	ALTER SEQUENCE accounts_id_seq RESTART WITH 1
	`
	_, errReset := conn.QueryContext(context.Background(), resetSequence)
	if errReset != nil {
		log.Fatal("errpr when reset sequence", err)
	}

}

func TestMain(m *testing.M) {
	config := LoadConfigDb()

	var err error
	conn, err = sql.Open("postgres", config)

	if err != nil {
		log.Fatal("cannot connect db due to: ", err)
	}
	defer conn.Close()

	truncateTableAccount()

	trx = NewTRX(conn)
	os.Exit(m.Run())
}
