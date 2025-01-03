package repository

import (
	"context"

	"taskservice/models"
)

type Repo interface {
	// Create
	CreateNew(ctx context.Context, record *models.Task) error

	// Read
	GetByTaskID(ctx context.Context, id uint64) (*models.Task, error)
	GetAllCreated(ctx context.Context, limit int, id uint64) ([]models.Task, error)
	GetAllAssigned(ctx context.Context, limit int, id uint64) ([]models.Task, error)

	// Update
	UpdateExisting(ctx context.Context, id uint64, record *models.Task) error

	// Delete
	DeleteByTaskID(ctx context.Context, id uint64) error
}
