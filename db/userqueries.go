package db

import (
	"ps_web_0/models"
	"github.com/vattle/sqlboiler/queries/qm"
)

/*
Returns users ordered alphabetically
 */
func GetAllUsers() (models.UserSlice, error) {
	return models.UsersG(qm.OrderBy("Name")).All()
}
func UpdateRole(userId float64, role int) (*models.User) {
	mUser, err := models.FindUserG(userId)
	if err != nil {
		panic(err)
	}

	mUser.Role = role
	mUser.UpdateG("role")

	return mUser
}