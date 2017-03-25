package main

import (
	"testing"
	"database/sql"
	"fmt"
	"github.com/vattle/sqlboiler/boil"
	"ps_web_0/models"
)

func TestUsers(*testing.T) {
	// Open handle to database like normal
	db, err := sql.Open("postgres", "dbname=coopcat user=coopcat_dev password=coopcat_dev")
	if err != nil {
		fmt.Print(err)
	}

	// If you don't want to pass in db to all generated methods
	// you can use boil.SetDB to set it globally, and then use
	// the G variant methods like so:
	boil.SetDB(db)

	users, err := models.UsersG().All()
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("OK")
		fmt.Println(len(users))
		for _, element := range users {
			fmt.Println(element)
		}
	}
}
