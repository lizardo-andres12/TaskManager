package repository

import "taskservice/models"

type Repo interface {
	// Create
	CreateNew(record *models.Task) error

	// Read
	GetByTaskID(id uint64) (*models.Task, error)
	GetByCreatorID(id uint64) (*models.Task, error)
	GetByAssigneeID(id uint64) (*models.Task, error)
	GetAllCreated(limit int, id uint64) ([]models.Task, error)
	GetAllAssigned(limit int, id uint64) ([]models.Task, error)

	// Update
	UpdateExisting(id uint64, record *models.Task) error

	// Delete
	DeleteByTaskID(id uint64) error
}
