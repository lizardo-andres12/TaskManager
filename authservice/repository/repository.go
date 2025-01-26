package repository

import (
	"context"

	"authservice/models"
)

type Repository interface {
	// Create
	CreateNewUserAuth(ctx context.Context, auth *models.UserAuth) error

	// Read
	GetUserDetails(ctx context.Context, id uint64) (*models.User, error)
	GetAuthByEmail(ctx context.Context, email string) (*models.Auth, error)
	GetAuthByUsername(ctx context.Context, username string) (*models.Auth, error)

	// Update
	UpdateUsername(ctx context.Context, id uint64, username string) error
	UpdatePassword(ctx context.Context, id uint64, password string, salt string) error
	// TODO: UpdateMediaURL once media storage is figured out

	// Delete
	DeleteUser(ctx context.Context, id uint64) error
}
