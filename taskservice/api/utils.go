package api

import (
	"taskservice/models"
	pb "taskservice/proto"
)

func TaskToPBTask(task *models.Task) *pb.Task {
	return &pb.Task{
		TaskId:      task.TaskID,
		Title:       task.Title,
		Description: task.Description,
		Status:      uint32(task.Status),
		Deadline:    task.Deadline,
		Priority:    task.Priority,
		CreatorId:   task.CreatorID,
		TeamId:      task.TeamID,
	}
}

func PBTaskToTask(task *pb.Task) *models.Task {
	return &models.Task{
		TaskID:      task.GetTaskId(),
		Title:       task.GetTitle(),
		Description: task.GetDescription(),
		Status:      uint8(task.GetStatus()),
		Deadline:    task.GetDeadline(),
		Priority:    task.GetPriority(),
		CreatorID:   task.GetCreatorId(),
		TeamID:      task.GetTeamId(),
	}
}
