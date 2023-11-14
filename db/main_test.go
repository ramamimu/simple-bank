package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
)

var trx TRX

func TestMain(m *testing.M) {
	config := LoadConfigDb()

	conn, err := sql.Open("postgres", config)

	if err != nil {
		log.Fatal("cannot connect db due to: ", err)
	}
	defer conn.Close()

	trx = NewTRX(conn)
	fmt.Println("-> ", trx)

	os.Exit(m.Run())
}
