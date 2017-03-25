package controller

import (
	"net/http"
	"ps_web_0/view"
	"ps_web_0/db"
	"fmt"
	"strconv"
)

const loggedInID  = "4"

func adminGET(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]interface{})

	users, _ := db.GetAllUsers()

	data["Users"] = users

	view.RenderTemplate(w, "admin", data)
}

func adminPOST(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]interface{})

	userIdInt, _ := strconv.Atoi(r.PostFormValue("UserId"))
	userId := float64(userIdInt)
	newRole, _ := strconv.Atoi(r.PostFormValue("role"))

	returnedUser := db.UpdateRole(userId, newRole)

	data["notification"] = fmt.Sprintf("Updated the role of %s", returnedUser.Name)

	users, _ := db.GetAllUsers()

	data["Users"] = users

	view.RenderTemplate(w, "admin", data)
}