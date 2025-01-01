package service

import (
	"context"
	"taskservice/models"
)

type Service interface {
	// Read functions
	GetByID(ctx context.Context, id uint64, idType string) (*models.Task, error)
	GetAll(ctx context.Context, userId uint64, userType uint8) ([]models.Task, error)

	// Create, Update, Delete functions
	CreateTask(ctx context.Context, task *models.Task) error
	UpdateTask(ctx context.Context, id uint64, task *models.Task) error
	DeleteTask(ctx context.Context, id uint64) error
}
