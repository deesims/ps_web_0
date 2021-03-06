package db

import (
	"github.com/deesims/ps_web_0/models"
	"github.com/vattle/sqlboiler/queries/qm"
	"gopkg.in/nullbio/null.v6"
	"time"
	"github.com/deesims/ps_web_0/util"
)

/*
Returns users ordered alphabetically
*/
func GetAllUsers() (models.UserSlice, error) {
	return models.UsersG(qm.OrderBy("Name")).All()
}
func UpdateRole(userId float64, role int) *models.User {
	mUser, err := models.FindUserG(userId)
	util.PanicOnError(err)
	mUser.Role = role

	mUser.UpdateGP("role")

	return mUser
}

func GetAllCompanies() (models.CompanySlice, error) {
	return models.CompaniesG(qm.OrderBy("Name")).All()
}

type JobWithCompanySlice []JobWithCompany

type JobWithCompany struct {
	JobID                 float64     `boil:"job_id" json:"job_id" toml:"job_id" yaml:"job_id"`
	CompanyID             float64     `boil:"company_id" json:"company_id" toml:"company_id" yaml:"company_id"`
	CompanyName           string      `boil:"companyname" json:"companyname" toml:"companyname" yaml:"companyname"`
	Name                  string `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Rating                null.Int    `boil:"rating" json:"rating,omitempty" toml:"rating" yaml:"rating,omitempty"`
	StudentLevel          int         `boil:"student_level" json:"student_level,omitempty" toml:"student_level" yaml:"student_level,omitempty"`
	NumOfPositions        int         `boil:"num_of_positions" json:"num_of_positions" toml:"num_of_positions" yaml:"num_of_positions"`
	DeadlineDate          time.Time   `boil:"deadline_date" json:"deadline_date" toml:"deadline_date" yaml:"deadline_date"`
	NumAvailablePositions int         `boil:"num_available_positions" json:"num_available_positions,omitempty" toml:"num_available_positions" yaml:"num_available_positions,omitempty"`
}
func (job *JobWithCompany) DeadlineDateFormatted() string {
	return job.DeadlineDate.Format("2006-02-01")
}

func GetAllJobsWithCompanyName() (JobWithCompanySlice, error) {
	jobWithCompany := make(JobWithCompanySlice, 0)
	error := models.NewQueryG(qm.Select("*"), qm.From("getAllJobsWithCompanyName()")).Bind(&jobWithCompany)
	return jobWithCompany, error
}
