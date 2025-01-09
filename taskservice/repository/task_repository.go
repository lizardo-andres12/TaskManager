// --TODO-- refactor all code to match new task schema
// use query: select * from tasks t join workers w on t.taskId = w.taskId where w.userId=?;
package repository

import (
	"context"
	"database/sql"

	"taskservice/models"
)

type TaskRepo struct {
	DB *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{DB: db}
}

func (taskrepo *TaskRepo) CreateNew(ctx context.Context, record *models.Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(ctx, "INSERT INTO tasks (Title, Priority, CreatorID, AssigneeID) VALUES (?, ?, ?, ?)", record.Title, record.Priority, record.CreatorID, record.AssigneeID)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) GetByTaskID(ctx context.Context, id uint64) (*models.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var task models.Task

	row := taskrepo.DB.QueryRowContext(ctx, "SELECT * FROM tasks WHERE taskId = ?", id)
	if err := row.Scan(&task.TaskID, &task.Title, &task.Priority, &task.CreatorID, &task.AssigneeID, &task.CreatedAt); err != nil {
		return nil, err
	}
	return &task, nil
}

func (taskrepo *TaskRepo) GetAllCreated(ctx context.Context, limit int, id uint64) ([]models.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var tasks []models.Task

	rows, err := taskrepo.DB.QueryContext(ctx, "SELECT * FROM tasks WHERE creatorId = ? LIMIT ?", id, limit)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.TaskID, &task.Title, &task.Priority, &task.CreatorID, &task.AssigneeID, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (taskrepo *TaskRepo) GetAllAssigned(ctx context.Context, limit int, id uint64) ([]models.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var tasks []models.Task

	rows, err := taskrepo.DB.QueryContext(ctx, "SELECT * FROM tasks WHERE assigneeId = ? LIMIT ?", id, limit)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.TaskID, &task.Title, &task.Priority, &task.CreatorID, &task.AssigneeID, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (taskrepo *TaskRepo) UpdateExisting(ctx context.Context, id uint64, record *models.Task) error { // this function should never receive taskId not stored
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(ctx, "UPDATE tasks SET title = ?, priority = ?, creatorId = ?, assigneeId = ? WHERE taskId = ?", record.Title, record.Priority, record.CreatorID, record.AssigneeID, id)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) DeleteByTaskID(ctx context.Context, id uint64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(ctx, "DELETE FROM tasks WHERE taskId = ?", id)
	if err != nil {
		return err
	}
	return nil
}
