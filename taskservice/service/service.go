package service

import (
	"context"
	"taskservice/models"
)

type Service interface {
	// Read functions
	GetByID(ctx context.Context, id uint64) (*models.Task, error)
	GetAll(ctx context.Context, userId uint64, userType bool) ([]models.Task, error)

	// Create, Update, Delete functions
	CreateTask(ctx context.Context, task *models.Task) error
	UpdateTask(ctx context.Context, id uint64, task *models.Task) error
	DeleteTask(ctx context.Context, id uint64) error
}
