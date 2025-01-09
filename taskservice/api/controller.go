package api

import (
	"context"
	"taskservice/models"
	"taskservice/service"
	"time"

	pb "taskservice/proto"
)

type TaskController struct {
	TaskService *service.TaskService
}

func NewTaskController(ts *service.TaskService) *TaskController {
	return &TaskController{
		TaskService: ts,
	}
}

func (tc *TaskController) CreateTask(ctx context.Context, req *pb.CreateRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.CreateTask(ctx, &models.Task{
		Title:      req.GetTitle(),
		Priority:   req.GetPriority(),
		CreatorID:  req.GetCreatorId(),
		AssigneeID: req.GetAssigneeId(),
	})

	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (tc *TaskController) GetTask(ctx context.Context, req *pb.IDOnlyRequest) (*pb.GetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	task, err := tc.TaskService.GetByID(ctx, req.GetId())
	if err != nil {
		return &pb.GetResponse{
			Success: false,
			Task:    &pb.Task{},
		}, err
	}
	return &pb.GetResponse{
		Success: true,
		Task:    TaskToPBTask(task),
	}, nil
}

func (tc *TaskController) GetAllTasks(ctx context.Context, req *pb.IDOnlyRequest) (*pb.GetAllResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tasks, err := tc.TaskService.GetAll(ctx, req.GetId(), true)
	if err != nil {
		return &pb.GetAllResponse{
			Success: false,
			Tasks:   []*pb.Task{TaskToPBTask(&models.Task{})},
		}, err
	}

	var pbTasks []*pb.Task
	for _, task := range tasks {
		pbTasks = append(pbTasks, TaskToPBTask(&task))
	}

	return &pb.GetAllResponse{
		Success: true,
		Tasks:   pbTasks,
	}, nil
}

func (tc *TaskController) UpdateTask(ctx context.Context, req *pb.UpdateRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.UpdateTask(ctx, PBTaskToTask(req.GetTask()))
	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (tc *TaskController) DeleteTask(ctx context.Context, req *pb.IDOnlyRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.DeleteTask(ctx, req.GetId())
	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}
