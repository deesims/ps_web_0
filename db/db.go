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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InsertUser() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var lastInsertId int
	db.Exec("INSERT INTO coopcat.user VALUES(13123, 'arnold this is from golang')")
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)

	db.Close()

	fmt.Println("Database closed.")
}
