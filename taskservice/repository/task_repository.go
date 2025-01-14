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

func (taskrepo *TaskRepo) CreateTask(ctx context.Context, task *models.Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(
		ctx,
		"INSERT INTO task (title, description, deadline, priority, creatorId, teamId) VALUES (?, ?, ?, ?, ?, ?, ?)",
		task.Title,
		task.Description,
		task.Status,
		task.Deadline,
		task.Priority,
		task.CreatorID,
		task.TeamID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) AssignToTask(ctx context.Context, taskAssignee *models.TaskAssignee) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(
		ctx,
		"INSERT INTO assignee (taskId, assigneeId) VALUES (?, ?)",
		taskAssignee.TaskID, taskAssignee.AssigneeID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) GetAllAssigned(ctx context.Context, id uint64, limit uint32, offset uint32) ([]models.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var tasks []models.Task

	rows, err := taskrepo.DB.QueryContext(
		ctx,
		"SELECT t.title, t.description, t.status, t.deadline, t.priority, t.creatorId, t.teamId FROM task t "+
			"INNER JOIN assignee a ON t.taskId = a.taskId WHERE a.assigneeId = ? ORDER BY t.createdAt DESC LIMIT ? OFFSET ?",
		id,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.Title, &task.Description, &task.Status, &task.Deadline, &task.Priority, &task.CreatorID, &task.TeamID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (taskrepo *TaskRepo) GetByTaskID(ctx context.Context, id uint64) (*models.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var task models.Task

	row := taskrepo.DB.QueryRowContext(
		ctx,
		"SELECT taskId, title, description, status, deadline, priority, creatorId, teamId FROM task WHERE taskId = ?",
		id,
	)

	if err := row.Scan(&task.Title, &task.Description, &task.Status, &task.Deadline, &task.Priority, &task.CreatorID, &task.TeamID); err != nil {
		return nil, err
	}
	return &task, nil
}

func (taskrepo *TaskRepo) UpdateStatus(ctx context.Context, id uint64, status uint8) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(ctx, "UPDATE task SET status = ? WHERE taskId = ?", status, id)
	if err != nil {
		return err
	}
	return nil
}
