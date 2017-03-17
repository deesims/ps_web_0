package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"os"
	"encoding/json"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "coopcat"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type ConnectionSettings struct {
	User string //TIL that variables starting with lowercase are private
	Password string
}

func InsertUser() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	conSettings := ConnectionSettings{}
	err := decoder.Decode(&conSettings)
	if err != nil {
		checkErr(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, conSettings.User, conSettings.Password, dbname)

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
