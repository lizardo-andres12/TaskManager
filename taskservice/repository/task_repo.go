package repository

import (
	"database/sql"
	"sync"
	"taskservice/models"
)

type TaskRepo struct {
	db *sql.DB
	mu sync.Mutex
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (taskrepo *TaskRepo) CreateNew(record *models.Task) error {
	taskrepo.mu.Lock()
	defer taskrepo.mu.Unlock()

	_, err := taskrepo.db.Exec("INSERT INTO tasks (Title, Priority, UserID, CreatedAt) VALUES (?, ?, ?, ?)", record.Title, record.Priority, record.UserID, record.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) GetByID(ids ...uint64) (*models.Task, error) {
	taskrepo.mu.Lock()
	defer taskrepo.mu.Unlock()
	var task models.Task

	row := taskrepo.db.QueryRow("SELECT * FROM tasks WHERE id = ?", ids[0])
	if err := row.Scan(&task.ID, &task.Title, &task.Priority, &task.UserID, &task.CreatedAt); err != nil {
		return nil, err
	}
	return &task, nil
}

func (taskrepo *TaskRepo) GetAll(limit int, id uint64) ([]models.Task, error) {
	taskrepo.mu.Lock()
	defer taskrepo.mu.Unlock()
	var tasks []models.Task

	rows, err := taskrepo.db.Query("SELECT * FROM tasks WHERE userId = ? LIMIT ?", id, limit)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Priority, &task.UserID, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (taskrepo *TaskRepo) UpdateExisting(id uint64, record *models.Task) error {
	taskrepo.mu.Lock()
	defer taskrepo.mu.Unlock()

	_, err := taskrepo.db.Exec("UPDATE tasks SET title = ?, priority = ? WHERE id = ?", record.Title, record.Priority, id)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) DeleteByID(id uint64) error {
	taskrepo.mu.Lock()
	defer taskrepo.mu.Unlock()

	_, err := taskrepo.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
