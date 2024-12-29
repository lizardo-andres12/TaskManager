// --TODO-- rework functions to match changes made to worker model
package repository

import (
	"database/sql"
	"taskservice/models"
)

// Deprecated: Worker is not part of database
type WorkerRepo struct {
	db *sql.DB
}

// Deprecated: Worker is not part of database
func NewWorkerRepo(db *sql.DB) *WorkerRepo {
	return &WorkerRepo{db: db}
}

// Deprecated: Worker is not part of database
func (workerrepo *WorkerRepo) CreateNew(record *models.Worker) error {
	_, err := workerrepo.db.Exec("INSERT INTO workers VALUES (userId, taskId, username)", record.UserID, record.TaskID, record.Username)
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: Worker is not part of database
func (workerrepo *WorkerRepo) GetByID(ids ...uint64) (*models.Worker, error) {
	var worker models.Worker

	row := workerrepo.db.QueryRow("SELECT * FROM workers WHERE userId = ? AND taskId = ?", ids[0], ids[1])
	if err := row.Scan(&worker.UserID, &worker.TaskID, &worker.Username); err != nil {
		return nil, err
	}
	return &worker, nil
}

// Deprecated: Worker is not part of database
func (workerrepo *WorkerRepo) GetAll(limit int, id uint64) ([]models.Worker, error) {
	var workers []models.Worker

	rows, err := workerrepo.db.Query("SELECT * FROM tasks WHERE userId = ? LIMIT ?", id, limit)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var worker models.Worker
		if err := rows.Scan(&worker.UserID, &worker.TaskID, &worker.Username); err != nil {
			return nil, err
		}
		workers = append(workers, worker)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return workers, nil
}

// Deprecated: Worker is not part of database
func (workerrepo *WorkerRepo) UpdateExisting(id uint64, record *models.Worker) error {
	// updating not needed as of now, must add some column to database that makes sense to update
	return nil
}

// Deprecated: Worker is not part of database
func (workerrepo *WorkerRepo) DeleteByID(ids ...uint64) error {
	_, err := workerrepo.db.Exec("DELETE FROM tasks WHERE userId = ? AND taskId = ?", ids[0], ids[1])
	if err != nil {
		return err
	}
	return nil
}
