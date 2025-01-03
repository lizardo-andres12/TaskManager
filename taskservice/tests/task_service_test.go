package tests

import (
	"database/sql"
	"testing"

	"taskservice/models"
	"taskservice/repository"
	"taskservice/service"
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

	ctx, cancel := loadContext()
	defer cancel()

	ts.DeleteTask(ctx, task)
}
