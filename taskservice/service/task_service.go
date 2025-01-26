package service

import (
	"context"
	"sync"
	"time"

	"taskservice/models"
	"taskservice/repository"
)

type TaskService struct {
	TaskRepository *repository.TaskRepository
}

func NewTaskService(tr *repository.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepository: tr,
	}
}

func (ts *TaskService) CreateTask(ctx context.Context, task *models.Task) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.CreateTask(ctx, task)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) AssignToTask(ctx context.Context, assignee *models.TaskAssignee) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.AssignToTask(ctx, assignee)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) GetAllAssigned(ctx context.Context, id uint64, limit uint64, offset uint64) ([]models.Task, error) {
	var wg sync.WaitGroup
	var tasks []models.Task
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		tasks, err = ts.TaskRepository.GetAllAssigned(ctx, id, limit, offset)
	}()
	wg.Wait()

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *TaskService) GetAllCreated(ctx context.Context, id uint64, limit uint64, offset uint64) ([]models.Task, error) {
	var wg sync.WaitGroup
	var tasks []models.Task
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		tasks, err = ts.TaskRepository.GetAllCreated(ctx, id, limit, offset)
	}()
	wg.Wait()

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *TaskService) GetByTaskID(ctx context.Context, id uint64) (*models.Task, error) {
	var wg sync.WaitGroup
	var task *models.Task
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		task, err = ts.TaskRepository.GetByTaskID(ctx, id)
	}()
	wg.Wait()

	if err != nil {
		return nil, err
	}
	return task, nil
}

func (ts *TaskService) UpdateTitle(ctx context.Context, id uint64, title string) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.UpdateTitle(ctx, id, title)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) UpdateDescription(ctx context.Context, id uint64, desc string) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.UpdateDescription(ctx, id, desc)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) UpdateStatus(ctx context.Context, id uint64, status uint8) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.UpdateStatus(ctx, id, status)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) UpdateDeadline(ctx context.Context, id uint64, deadline *time.Time) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.UpdateDeadline(ctx, id, deadline)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) UpdatePriority(ctx context.Context, id uint64, priority bool) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.UpdatePriority(ctx, id, priority)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) DeleteTask(ctx context.Context, id uint64) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.DeleteTask(ctx, id)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) UnassignTask(ctx context.Context, id uint64) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.UnassignTask(ctx, id)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}
