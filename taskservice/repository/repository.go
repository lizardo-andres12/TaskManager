package repository

import (
	"context"
	"time"

	"taskservice/models"
)

type Repository interface {
	// Create
	CreateTask(ctx context.Context, task *models.Task) error
	AssignToTask(ctx context.Context, taskAssignee *models.TaskAssignee) error

	// Read
	GetAllAssigned(ctx context.Context, id uint64, limit uint64, offset uint64) ([]models.Task, error)
	GetAllCreated(ctx context.Context, creatorId uint64, limit uint64, offset uint64) ([]models.Task, error)
	GetByTaskID(ctx context.Context, id uint64) (*models.Task, error)

	// Update
	UpdateTitle(ctx context.Context, id uint64, title string) error
	UpdateDescription(ctx context.Context, id uint64, description string) error
	UpdateStatus(ctx context.Context, id uint64, status uint8) error
	UpdateDeadline(ctx context.Context, id uint64, deadline *time.Time) error
	UpdatePriority(ctx context.Context, id uint64, priority bool) error

	// Delete
	DeleteTask(ctx context.Context, id uint64) error
	UnassignTask(ctx context.Context, id uint64) error
}
