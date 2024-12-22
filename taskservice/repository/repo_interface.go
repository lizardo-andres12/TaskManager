package repository

import "taskservice/models"

type Repo[T models.Record] interface {
	// Create
	CreateNew(record *T) error

	// Read
	GetByID(id uint64) (*T, error)
	GetAll(limit int, ids ...uint64) ([]T, error)

	// Update
	UpdateExisting(id uint64, record *T) error

	// Delete
	DeleteByID(id uint64)
}
