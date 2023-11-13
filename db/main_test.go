package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var db Db

func TestMain(m *testing.M) {
	config := LoadConfigDb()

	conn, err := sql.Open("postgres", config)

	if err != nil {
		log.Fatal("cannot connect db due to: ", err)
	}
	defer conn.Close()

	db = NewDb(conn)

	os.Exit(m.Run())
}
