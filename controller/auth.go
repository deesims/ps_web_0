package controller

import (
	"fmt"
	"net/http"

	"github.com/deesims/ps_web_0/db"

	"github.com/apexskier/httpauth"
)

var (
	authHandler httpauth.Authorizer
)

func AuthInit(w http.ResponseWriter, r *http.Request) {
	dbAuth, err := httpauth.NewSqlAuthBackend("postgres", db.PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	roles := make(map[string]httpauth.Role)
	roles["user"] = 30
	roles["admin"] = 80

	authHandler, err = httpauth.NewAuthorizer(dbAuth, []byte("cookie-encryption-key"), "user", roles)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	username := "testUser"

	defaultUser := httpauth.UserData{Username: username, Role: "admin"}
	err = dbAuth.SaveUser(defaultUser)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	// Update user with a password and email address
	err = authHandler.Update(w, r, username, "adminadmin", "admin@localhost.com")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
