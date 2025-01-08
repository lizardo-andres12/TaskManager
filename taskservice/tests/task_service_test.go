package tests

import (
	"database/sql"
	"testing"

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

	ctx, cancel := loadContext()
	defer cancel()

	ts.DeleteTask(ctx, 1)
}
