package repository

import (
	"database/sql"
	"loginservice/models"
)

type UserRepo struct {
	DB *sql.DB
}

func (userRepo *UserRepo) CreateNewUser(user *models.User) error {

	query := "INSERT INTO users (name, email, password, manager) VALUES (?, ?, ?, ?)"
	_, err := userRepo.DB.Exec(query, user.Name, user.Email, user.Password, user.Manager)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *UserRepo) UpdateUser(updatedField string, updatedvalue any, userID uint64) error {

	_, err := userRepo.DB.Exec("UPDATE users SET ? = ? WHERE id = ? ", updatedField, updatedvalue, userID)
	if err != nil {
		return err
	}
	return nil
}
