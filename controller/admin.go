package controller

import (
	"fmt"
	"github.com/deesims/ps_web_0/db"
	"github.com/deesims/ps_web_0/models"
	"github.com/deesims/ps_web_0/view"
	"github.com/gorilla/schema"
	"net/http"
	"strconv"
	"time"
)

func checkAdminRole(w http.ResponseWriter, r *http.Request) bool {
	currentUser, err := authHandler.CurrentUser(w, r)
	if err != nil {
		fmt.Println("Error getting user: ", err.Error())
		return false
	}
	if currentUser.Role != "admin" {
		fmt.Println("Current user is not an admin")
		return false
	}

	return true
}

func adminRoles(w http.ResponseWriter, r *http.Request) {

	if !checkAdminRole(w, r) {
		fmt.Println("Error, not an admin.")
		return
	}

	var data = make(map[string]interface{})

	if r.Method == "POST" {
		userIDInt, _ := strconv.Atoi(r.PostFormValue("UserId"))
		userID := float64(userIDInt)
		newRole, _ := strconv.Atoi(r.PostFormValue("role"))

		returnedUser := db.UpdateRole(userID, newRole)

		data["notification"] = fmt.Sprintf("Updated the role of %s", returnedUser.Name)
	}

	users, _ := db.GetAllUsers()
	data["Users"] = users

	view.RenderTemplate(w, "admin_roles", data)
}

func adminAddJob(w http.ResponseWriter, r *http.Request) {

	if !checkAdminRole(w, r) {
		fmt.Println("Error, not an admin.")
		return
	}

	var data = make(map[string]interface{})

	if r.Method == "POST" {
		returnedJob := new(models.Job)
		r.ParseForm() //Need to call before r.PostForm
		schema.NewDecoder().Decode(returnedJob, r.PostForm)

		//Have to manually parse date
		deadlineDate, _ := time.Parse("2006-02-01", r.FormValue("DeadlineDate"))
		returnedJob.DeadlineDate = deadlineDate

		if returnedJob.JobID == 0 {
			returnedJob.InsertG()
		} else {
			returnedJob.UpdateG()
		}

		data["notification"] = "Updated"
	}
	//To select the company from a list when adding a job
	companies, _ := db.GetAllCompanies()
	data["Companies"] = companies

	jobs, error := db.GetAllJobsWithCompanyName()
	if error != nil {
		panic(error)
	}

	data["Jobs"] = jobs

	view.RenderTemplate(w, "admin_addjob", data)
}

func adminCompanies(w http.ResponseWriter, r *http.Request) {

	if !checkAdminRole(w, r) {
		fmt.Println("Error, not an admin.")
		return
	}

	var data = make(map[string]interface{})

	if r.Method == "POST" {
		returnedCompany := new(models.Company)
		r.ParseForm() //Need to call before r.PostForm
		schema.NewDecoder().Decode(returnedCompany, r.PostForm)

		if returnedCompany.CompanyID == 0 {
			returnedCompany.InsertG()
		} else {
			returnedCompany.UpdateG()
		}

		data["notification"] = "Updated"
	}

	companies, _ := db.GetAllCompanies()
	data["Companies"] = companies

	view.RenderTemplate(w, "admin_companies", data)
}
