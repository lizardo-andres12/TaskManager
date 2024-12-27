package repository

import "taskservice/models"

type Repo[T models.Record] interface {
	// Create
	CreateNew(record *T) error

	// Read
	GetByID(ids ...uint64) (*T, error) // variadic because of dual primary key record
	GetAll(limit int, id uint64) ([]T, error)

	// Update
	UpdateExisting(id uint64, record *T) error

	// Delete
	DeleteByID(ids ...uint64) error // variadic because of dual primary key record
}
