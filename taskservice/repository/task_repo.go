// --TODO-- implement method GetAllAssigned
// use query: select * from tasks t join workers w on t.taskId = w.taskId where w.userId=?;
package repository

import (
	"database/sql"
	"taskservice/repository/models"
)

type TaskRepo struct {
	DB *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{DB: db}
}

func (taskrepo *TaskRepo) CreateNew(record *models.Task) error {
	_, err := taskrepo.DB.Exec("INSERT INTO tasks (Title, Priority, UserID, CreatedAt) VALUES (?, ?, ?, ?)", record.Title, record.Priority, record.UserID, record.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) GetByID(ids ...uint64) (*models.Task, error) {
	var task models.Task

	row := taskrepo.DB.QueryRow("SELECT * FROM tasks WHERE taskId = ?", ids[0])
	if err := row.Scan(&task.TaskID, &task.Title, &task.Priority, &task.UserID, &task.CreatedAt); err != nil {
		return nil, err
	}
	return &task, nil
}

// Gets all tasks CREATED by userId
func (taskrepo *TaskRepo) GetAll(limit int, id uint64) ([]models.Task, error) {
	var tasks []models.Task

	rows, err := taskrepo.DB.Query("SELECT * FROM tasks WHERE userId = ? LIMIT ?", id, limit)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.TaskID, &task.Title, &task.Priority, &task.UserID, &task.CreatedAt); err != nil {
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
	_, err := taskrepo.DB.Exec("UPDATE tasks SET title = ?, priority = ?, userId = ? WHERE taskId = ?", record.Title, record.Priority, record.UserID, id)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) DeleteByID(id uint64) error {
	_, err := taskrepo.DB.Exec("DELETE FROM tasks WHERE taskId = ?", id)
	if err != nil {
		return err
	}
	return nil
}
