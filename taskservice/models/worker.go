// --TODO-- add createdAt and completed fields
package models

type Worker struct {
	UserID   uint64 `json:"userId"` // Primary Key/ForeignKey
	TaskID   uint64 `json:"taskId"` // Primary Key
	Username string `json:"username"`
}
