package models

import "time"

type Task struct {
	ID        uint64    `json:"id"` // Primary Key
	Title     string    `json:"title"`
	Priority  bool      `json:"priority"`
	UserID    uint64    `json:"userId"` // Foreign Key
	CreatedAt time.Time `json:"createdAt"`
}
