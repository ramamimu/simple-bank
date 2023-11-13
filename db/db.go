package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "postgres"
	dbname   string = "simplebank"
)

func LoadConfigDb() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

type TRX interface {
	CreateAccount() Account
}

type Db struct {
	conn *sql.DB
}

func NewDb(conn *sql.DB) Db {
	return Db{
		conn: conn,
	}

}
