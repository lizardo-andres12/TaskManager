package models

import "time"

type Task struct {
	TaskID      uint64    `json:"taskId"` // Primary Key
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      uint8     `json:"status"`
	Deadline    time.Time `json:"deadline"`
	Priority    bool      `json:"priority"`
	CreatorID   uint64    `json:"creatorId"` // Foreign Key
	TeamID      uint64    `json:"teamId"`    // Nullable
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
