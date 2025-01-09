package service

import (
	"context"
	"fmt"
	"sync"

	"taskservice/models"
	"taskservice/repository"
)

type TaskService struct {
	TaskRepository *repository.TaskRepo
}

func NewTaskService(tr *repository.TaskRepo) *TaskService {
	return &TaskService{TaskRepository: tr}
}

func (ts *TaskService) GetByID(ctx context.Context, id uint64) (*models.Task, error) {
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

func (ts *TaskService) GetAll(ctx context.Context, userId uint64, userType bool) ([]models.Task, error) {
	var wg sync.WaitGroup
	var tasks []models.Task
	var err error
	wg.Add(1)

	if userType {
		go func() {
			defer wg.Done()
			tasks, err = ts.TaskRepository.GetAllCreated(ctx, 10000, userId)
		}()
	} else {
		go func() {
			defer wg.Done()
			tasks, err = ts.TaskRepository.GetAllAssigned(ctx, 10000, userId)
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
		err = ts.TaskRepository.CreateNew(ctx, task)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) UpdateTask(ctx context.Context, taskId uint64, task *models.Task) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	fmt.Println(taskId)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.UpdateExisting(ctx, taskId, task)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) DeleteTask(ctx context.Context, taskId uint64) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepository.DeleteByTaskID(ctx, taskId)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}
