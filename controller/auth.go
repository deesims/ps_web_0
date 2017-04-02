package controller

import (
	"fmt"
	"net/http"

	"github.com/deesims/ps_web_0/db"
	"github.com/deesims/ps_web_0/models"

	"github.com/apexskier/httpauth"
)

var (
	authHandler httpauth.Authorizer
)

// AuthInit initializes the authentication system for logging in and session storing
func AuthInit(w http.ResponseWriter, r *http.Request) {
	dbAuth, err := httpauth.NewSqlAuthBackend("postgres", db.PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	roles := make(map[string]httpauth.Role)
	roles["user"] = 30
	roles["moderator"] = 60
	roles["admin"] = 80

	authHandler, err = httpauth.NewAuthorizer(dbAuth, []byte("cookie-encryption-key"), "user", roles)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	adminName := "admin"
	moderatorName := "moderator"
	userName := "user"

	admin := httpauth.UserData{Username: adminName, Role: "admin"}
	err = dbAuth.SaveUser(admin)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	// Update user with a password and email address
	err = authHandler.Update(w, r, adminName, adminName, "admin@localhost.com")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	moderator := httpauth.UserData{Username: moderatorName, Role: "moderator"}
	err = dbAuth.SaveUser(moderator)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = authHandler.Update(w, r, moderatorName, moderatorName, "moderator@localhost.com")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	user := httpauth.UserData{Username: userName, Role: "user"}
	err = dbAuth.SaveUser(user)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = authHandler.Update(w, r, userName, userName, "user@localhost.com")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	adminDB := new(models.User)
	adminDB.Name = adminName
	adminDB.Role = 1
	adminDB.Email = "admin@localhost.com"
	if err = adminDB.InsertG(); err != nil {
		fmt.Println("Error inserting admin...", err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	moderatorDB := new(models.User)
	moderatorDB.Name = moderatorName
	moderatorDB.Role = 2
	moderatorDB.Email = "moderator@localhost.com"
	if err = moderatorDB.InsertG(); err != nil {
		fmt.Println("Error inserting moderator...", err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	userDB := new(models.User)
	userDB.Name = userName
	userDB.Role = 3
	userDB.Email = "user@localhost.com"
	if err = userDB.InsertG(); err != nil {
		fmt.Println("Error inserting...", err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
