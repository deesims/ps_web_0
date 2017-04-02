package controller

import (
	"fmt"
	"github.com/deesims/ps_web_0/db"
	"github.com/deesims/ps_web_0/models"
	"github.com/deesims/ps_web_0/view"
	"gopkg.in/nullbio/null.v6"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func registerRoutesToFuncs(r *mux.Router) {
	r.HandleFunc("/auth", AuthInit)
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/admin/roles", adminRoles).Methods("GET", "POST")
	r.HandleFunc("/admin/jobs", adminJobs).Methods("GET", "POST")
	r.HandleFunc("/admin/companies", adminCompanies).Methods("GET", "POST")
	r.HandleFunc("/moderator", moderatorResumeSummary).Methods("GET", "POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/userhub", GetUserHubHandler).Methods("GET")
	r.HandleFunc("/sendresumetomod", SendResumeToModerator).Methods("POST")
	r.HandleFunc("/viewresume", ViewResume).Methods("GET")
	r.HandleFunc("/checkUser", checkUser)
}

// ViewResume gets the resume of logged in user
func ViewResume(w http.ResponseWriter, r *http.Request) {
	currentUser, err := authHandler.CurrentUser(w, r)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	user := db.FindUserFromUsername(currentUser.Username)
	resumes := db.FindAllResumesForAuthorID(user.UserID)

	data := map[string]interface{}{
		"CurrentUser": currentUser,
		"Resumes":     resumes,
	}

	view.RenderTemplate(w, "viewresume", data)
}

func LoadFileToDB(w http.ResponseWriter, r *http.Request) {
	return
}

func SendResumeToModerator(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer file.Close()

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Getwd() error: ", err.Error())
	}

	fmt.Println("current working dir:", currentDir)

	resumeDir := "/public/resumes/"

	out, err := os.Create(currentDir + resumeDir + header.Filename)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	var filePath null.String
	filePath.SetValid(resumeDir + header.Filename)

	currentUser, err := authHandler.CurrentUser(w, r)
	if err != nil {
		return
	}

	author := db.FindUserFromUsername(currentUser.Username)

	fmt.Println("author of the resume", author.Name)

	resumeObject := models.Resume{
		AuthorID:   (float64)(author.UserID),
		ResumePath: filePath,
	}

	err = resumeObject.InsertG()
	if err != nil {
		fmt.Println("resume object insertg error:", err.Error())
	}

	http.Redirect(w, r, "/viewresume", http.StatusSeeOther)
}

func GetUserHubHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserHub Executing...")
	currentUser, err := authHandler.CurrentUser(w, r)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	if currentUser.Role != "user" {
		log.Print("Error: Not a user")
		return
	}

	data := map[string]interface{}{
		"CurrentUser": currentUser,
	}

	view.RenderTemplate(w, "userhub", data)

}

// Handles the index page, renders a home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homeHandler Executing...")
	data := map[string]interface{}{
		"hello-user":  103,
		"thats-rught": 104,
	}
	view.RenderTemplate(w, "index", data)
}

// loginHandler
func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginHandler Executing...")
	currentUser, err := authHandler.CurrentUser(w, r)
	if err != nil {
		log.Println("Error: ", err.Error())
	}

	data := map[string]interface{}{
		"LoggedUser": currentUser,
		"lol":        "103",
	}

	view.RenderTemplate(w, "login", data)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}

	username := r.FormValue("lg_username")
	password := r.FormValue("lg_password")

	authHandler.Login(w, r, username, password, "/checkUser") // catch seem to catch proper err for redirect

}

func checkUser(w http.ResponseWriter, r *http.Request) {
	user, err := authHandler.CurrentUser(w, r) // check what user logged in
	if err != nil {
		fmt.Println("user err at 156: ", err.Error())
		return
	}
	fmt.Println("Role of user: ", user.Role)

	if user.Role == "user" {
		http.Redirect(w, r, "/userhub", http.StatusSeeOther)
	} else if user.Role == "admin" {
		fmt.Println("executing admin add job page")
		http.Redirect(w, r, "/admin/addjob", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/moderator", http.StatusSeeOther)
	}
}

// Init initializes the controller and registers the routes to appropriate
// function handlers
func Init() {

	router := mux.NewRouter()
	registerRoutesToFuncs(router)

	http.Handle("/", router)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("Public directory launched;")
}
