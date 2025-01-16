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
	pb.UnimplementedTaskServiceGRPCServer
}

func NewTaskController(ts *service.TaskService) *TaskController {
	return &TaskController{
		TaskService: ts,
	}
}

func (tc *TaskController) CreateTask(ctx context.Context, req *pb.CreateRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.CreateTask(ctx, PBTaskToTask(req.GetTask()))
	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (tc *TaskController) AssignToTask(ctx context.Context, req *pb.AssignRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.AssignToTask(ctx, &models.TaskAssignee{
		TaskID:     req.GetTaskId(),
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

	task, err := tc.TaskService.GetByTaskID(ctx, req.GetId())
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

func (tc *TaskController) GetAllAssigned(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tasks, err := tc.TaskService.GetAllAssigned(ctx, req.GetId(), req.GetLimit(), req.GetOffset())
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

func (tc *TaskController) GetAllCreated(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tasks, err := tc.TaskService.GetAllCreated(ctx, req.GetId(), req.GetLimit(), req.GetOffset())
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

func (tc *TaskController) UpdateTitle(ctx context.Context, req *pb.UpdateStringRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.UpdateTitle(ctx, req.GetTaskId(), req.GetText())
	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (tc *TaskController) UpdateDescription(ctx context.Context, req *pb.UpdateStringRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.UpdateDescription(ctx, req.GetTaskId(), req.GetText())
	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (tc *TaskController) UpdateStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.UpdateStatus(ctx, req.GetTaskId(), uint8(req.GetStatus()))
	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (tc *TaskController) UpdateDeadline(ctx context.Context, req *pb.UpdateStringRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	deadline, err := time.Parse(time.DateTime, req.GetText())
	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}

	err = tc.TaskService.UpdateDeadline(ctx, req.GetTaskId(), &deadline)
	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (tc *TaskController) UpdatePriority(ctx context.Context, req *pb.UpdatePriorityRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.UpdatePriority(ctx, req.GetTaskId(), req.GetPriority())
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

func (tc *TaskController) UnassignTask(ctx context.Context, req *pb.IDOnlyRequest) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := tc.TaskService.UnassignTask(ctx, req.GetId())
	if err != nil {
		return &pb.SuccessResponse{
			Success: false,
		}, err
	}
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}
