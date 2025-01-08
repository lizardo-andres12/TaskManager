package controller

import (
	"taskservice/models"
	pb "taskservice/proto"
)

func TaskToPBTask(task *models.Task) *pb.Task {
	return &pb.Task{
		TaskId:     task.TaskID,
		Title:      task.Title,
		Priority:   task.Priority,
		CreatorId:  task.CreatorID,
		AssigneeId: task.AssigneeID,
		CreatedAt:  task.CreatedAt,
	}
}

func PBTaskToTask(task *pb.Task) *models.Task {
	return &models.Task{
		TaskID:     task.GetTaskId(),
		Title:      task.GetTitle(),
		Priority:   task.GetPriority(),
		CreatorID:  task.GetCreatorId(),
		AssigneeID: task.GetAssigneeId(),
		CreatedAt:  task.GetCreatedAt(),
	}
}
