package models

type Task struct {
	TaskID     uint64 `json:"taskId"` // Primary Key
	Title      string `json:"title"`
	Priority   bool   `json:"priority"`
	CreatorID  uint64 `json:"creatorId"`
	AssigneeID uint64 `json:"assigneeId"`
	CreatedAt  string `json:"createdAt"`
}
