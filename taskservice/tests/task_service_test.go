package tests

import (
	"context"
	"database/sql"
	"taskservice/models"
	"taskservice/repository"
	"taskservice/service"
	"testing"
	"time"
)

func TestDeleteTask(t *testing.T) {
	var db *sql.DB
	db, strErr := loadDB()
	if strErr != "" {
		t.Error(strErr)
	}
	defer db.Close()

	tr := repository.NewTaskRepo(db)
	ts := service.NewTaskService(tr)
	task := &models.Task{TaskID: 1}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ts.DeleteTask(ctx, task)
}
