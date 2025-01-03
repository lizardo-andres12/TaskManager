package service

import (
	"context"
	"sync"

	"taskservice/models"
	"taskservice/repository"
)

type TaskService struct {
	TaskRepo repository.Repo
}

func NewTaskService(tr repository.Repo) *TaskService {
	return &TaskService{TaskRepo: tr}
}

func (ts *TaskService) GetByID(ctx context.Context, id uint64) (*models.Task, error) {
	var wg sync.WaitGroup
	var task *models.Task
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		task, err = ts.TaskRepo.GetByTaskID(ctx, id)
	}()
	wg.Wait()

	if err != nil {
		return nil, err
	}
	return task, nil
}

func (ts *TaskService) GetAll(ctx context.Context, userId uint64, userType bool) ([]models.Task, error) {
	var wg sync.WaitGroup
	var tasks []models.Task
	var err error
	wg.Add(1)

	if userType {
		go func() {
			defer wg.Done()
			tasks, err = ts.TaskRepo.GetAllCreated(ctx, 10000, userId)
		}()
	} else {
		go func() {
			defer wg.Done()
			tasks, err = ts.TaskRepo.GetAllAssigned(ctx, 10000, userId)
		}()
	}
	wg.Wait()

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *TaskService) CreateTask(ctx context.Context, task *models.Task) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepo.CreateNew(ctx, task)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) UpdateTask(ctx context.Context, task *models.Task) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepo.UpdateExisting(ctx, task.TaskID, task)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) DeleteTask(ctx context.Context, task *models.Task) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepo.DeleteByTaskID(ctx, task.TaskID)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}
