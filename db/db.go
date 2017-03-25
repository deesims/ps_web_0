package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vattle/sqlboiler/boil"
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

func init()  {
	// Open handle to database like normal
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Print(err)
	}

	// If you don't want to pass in db to all generated methods
	// you can use boil.SetDB to set it globally, and then use
	// the G variant methods like so:
	boil.SetDB(db)
}