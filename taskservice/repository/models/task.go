package models

type Task struct {
	TaskID    uint64 `json:"taskId"` // Primary Key
	Title     string `json:"title"`
	Priority  bool   `json:"priority"`
	UserID    uint64 `json:"userId"` // Foreign Key
	CreatedAt string `json:"createdAt"`
}
