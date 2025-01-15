package models

type Task struct {
	TaskID      uint64 `json:"taskId"` // Primary Key
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      uint8  `json:"status"`
	Deadline    string `json:"deadline"`
	Priority    bool   `json:"priority"`
	CreatorID   uint64 `json:"creatorId"` // Foreign Key
	TeamID      uint64 `json:"teamId"`    // Nullable
}
