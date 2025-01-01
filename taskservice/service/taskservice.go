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

func (ts *TaskService) CreateTask(ctx context.Context, task *models.Task) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = ts.TaskRepo.CreateNew(task)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}
