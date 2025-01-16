package repository

import (
	"context"
	"database/sql"
	"time"

	"taskservice/models"
)

type TaskRepo struct {
	DB *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{
		DB: db,
	}
}

func (taskrepo *TaskRepo) CreateTask(ctx context.Context, task *models.Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(
		ctx,
		"INSERT INTO task (title, description, status, deadline, priority, creatorId, teamId) VALUES (?, ?, ?, ?, ?, ?, ?)",
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
		taskAssignee.TaskID,
		taskAssignee.AssigneeID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) GetAllAssigned(ctx context.Context, id uint64, limit uint64, offset uint64) ([]models.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var tasks []models.Task

	rows, err := taskrepo.DB.QueryContext(
		ctx,
		"SELECT t.taskId, t.title, t.description, t.status, t.deadline, t.priority, t.creatorId, IFNULL(t.teamId, 0) FROM task t "+
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
		err = rows.Scan(&task.TaskID, &task.Title, &task.Description, &task.Status, &task.Deadline, &task.Priority, &task.CreatorID, &task.TeamID)
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

func (taskrepo *TaskRepo) GetAllCreated(ctx context.Context, creatorId uint64, limit uint64, offset uint64) ([]models.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var tasks []models.Task

	rows, err := taskrepo.DB.QueryContext(
		ctx,
		"SELECT taskId, title, description, status, deadline, priority, creatorId, IFNULL(teamId, 0) FROM task "+
			"WHERE creatorId = ? ORDER BY createdAt DESC LIMIT ? OFFSET ?",
		creatorId,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.TaskID, &task.Title, &task.Description, &task.Status, &task.Deadline, &task.Priority, &task.CreatorID, &task.TeamID)
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
		"SELECT taskId, title, description, status, deadline, priority, creatorId, IFNULL(teamId, 0) FROM task WHERE taskId = ?",
		id,
	)

	if err := row.Scan(&task.TaskID, &task.Title, &task.Description, &task.Status, &task.Deadline, &task.Priority, &task.CreatorID, &task.TeamID); err != nil {
		return nil, err
	}
	return &task, nil
}

func (taskrepo *TaskRepo) UpdateTitle(ctx context.Context, id uint64, title string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(ctx, "UPDATE task SET title = ? WHERE taskId = ?", title, id)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) UpdateDescription(ctx context.Context, id uint64, description string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(ctx, "UPDATE task SET description = ? WHERE taskId = ?", description, id)
	if err != nil {
		return err
	}
	return nil
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

func (taskrepo *TaskRepo) UpdateDeadline(ctx context.Context, id uint64, deadline *time.Time) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	deadlineStr := deadline.Format(time.DateTime)

	_, err := taskrepo.DB.ExecContext(ctx, "UPDATE task SET deadline = ? WHERE taskId = ?", deadlineStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) UpdatePriority(ctx context.Context, id uint64, priority bool) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(ctx, "UPDATE task SET priority = ? WHERE taskId = ?", priority, id)
	if err != nil {
		return err
	}
	return nil
}

// TODO: implement cascading delete to assignee table
func (taskrepo *TaskRepo) DeleteTask(ctx context.Context, id uint64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(ctx, "DELETE FROM task WHERE taskId = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (taskrepo *TaskRepo) UnassignTask(ctx context.Context, id uint64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := taskrepo.DB.ExecContext(ctx, "DELETE FROM assignee WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
