package controller

import (
	"fmt"
	"github.com/deesims/ps_web_0/models"
	"github.com/deesims/ps_web_0/view"
	"gopkg.in/nullbio/null.v6"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func registerRoutesToFuncs(r *mux.Router) {
	r.HandleFunc("/", loginHandler)

	r.HandleFunc("/admin/roles", adminRoles).Methods("GET", "POST")
	r.HandleFunc("/admin/addjob", adminJobs).Methods("GET", "POST")
	r.HandleFunc("/admin/companies", adminCompanies).Methods("GET", "POST")

	r.HandleFunc("/moderator", moderatorResumeSummary).Methods("GET", "POST")

	r.HandleFunc("/login", loginHandler).Methods("GET", "POST")

	r.HandleFunc("/userhub", GetUserHubHandler).Methods("GET")
	r.HandleFunc("/sendresumetomod", SendResumeToModerator).Methods("POST")
	r.HandleFunc("/viewresume", ViewResume).Methods("GET")
}

func ViewResume(w http.ResponseWriter, r *http.Request) {
	currentUser, err := authHandler.CurrentUser(w, r)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	filePath := "public/resumes/resume.pdf"

	data := map[string]interface{}{
		"CurrentUser": currentUser,
		"File-Path":   filePath,
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

	resumeObject := models.Resume{
		AuthorID:   4,
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

	data := map[string]interface{}{
		"CurrentUser": currentUser,
	}

	view.RenderTemplate(w, "userhub", data)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err.Error())
		}

		username := r.FormValue("lg_username")
		password := r.FormValue("lg_password")

		if err = authHandler.Login(w, r, username, password, "/"); err != nil {
			http.Redirect(w, r, "/userhub", http.StatusSeeOther)
		}
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	} else {
	view.RenderTemplate(w, "login", nil)
	}
}

func Init() {

	router := mux.NewRouter()
	registerRoutesToFuncs(router)

	http.Handle("/", router)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("Public directory launched;")
}
