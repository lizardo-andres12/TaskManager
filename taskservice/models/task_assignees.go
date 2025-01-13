package models

type TaskAssignee struct {
	ID         uint64 `json:"id"`         // Primary Key
	TaskID     uint64 `json:"taskId"`     // Foreign Key -> Tasks
	AssigneeID uint64 `json:"assigneeId"` // Foreign Key - Users
}
