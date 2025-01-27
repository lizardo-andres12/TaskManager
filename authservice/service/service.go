package service

import (
	"authservice/models"
	"context"
)

type Service interface {
	// Creates new user with all data aggregated from register options on front end
	Register(ctx context.Context, user *models.UserAuth) error

	// Should first use a GetAuth repository method to get password information and id. Upon
	// successful authorization, GetUserDetails is called and a pointer to a User object is returned.
	// JWT must also be created within this method and returned
	Login(ctx context.Context, indexType string, index string, password string) (*models.User, string, error)

	// Directly unauthorizes JWT
	Logout(ctx context.Context, id uint64) error

	UpdateUsername(ctx context.Context, id uint64, username string) error
	UpdatePassword(ctx context.Context, id uint64, password string) error

	DeleteUser(ctx context.Context, id uint64) error
}
