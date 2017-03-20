package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "coopcat_dev"
	password = "coopcat_dev"
	dbname   = "coopcat"
)

var PsqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Connection() *sql.DB {
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}
