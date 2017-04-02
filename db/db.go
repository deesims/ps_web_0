package db

import (
	"database/sql"
	"fmt"
	"github.com/deesims/ps_web_0/models"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "coopcat"
)

var PsqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func init() {
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

func FindAllResumesForAuthorID(authorID float64) []models.Resume {
	var resumes []models.Resume

	authorid := strconv.FormatFloat(authorID, 'f', 6, 64)

	queries.RawG("SELECT * FROM resume " +
		"WHERE author_id=" + authorid).BindP(&resumes)

	return resumes
}

func FindUserFromUsername(username string) models.User {
	var user models.User
	queries.RawG("SELECT * FROM users WHERE ( name = " + username + " )").BindP(&user)
	return user
}
