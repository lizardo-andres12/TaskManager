package repository

import "loginservice/models"

type Repo interface {

	//create
	CreateNewUser(user *models.User) error

	//update
	UpdateUser(updateField string, updatedValue any, userID uint64)
}
