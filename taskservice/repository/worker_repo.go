package repository

import (
	"database/sql"
	"sync"
	"taskservice/models"
)

type WorkerRepo struct {
	db *sql.DB
	mu sync.Mutex
}

func NewWorkerRepo(db *sql.DB) *WorkerRepo {
	return &WorkerRepo{db: db}
}

func (workerrepo *WorkerRepo) CreateNew(record *models.Worker) error {
	workerrepo.mu.Lock()
	defer workerrepo.mu.Unlock()

	_, err := workerrepo.db.Exec("INSERT INTO workers VALUES (userId, taskId, username)", record.UserID, record.TaskID, record.Username)
	if err != nil {
		return err
	}
	return nil
}

func (workerrepo *WorkerRepo) GetByID(ids ...uint64) (*models.Worker, error) {
	workerrepo.mu.Lock()
	defer workerrepo.mu.Unlock()
	var worker models.Worker

	row := workerrepo.db.QueryRow("SELECT * FROM workers WHERE userId = ? AND taskId = ?", ids[0], ids[1])
	if err := row.Scan(&worker.UserID, &worker.TaskID, &worker.Username); err != nil {
		return nil, err
	}
	return &worker, nil
}

func (workerrepo *WorkerRepo) GetAll(limit int, id uint64) ([]models.Worker, error) {
	workerrepo.mu.Lock()
	defer workerrepo.mu.Unlock()
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
