// --TODO-- rewrite test functions involving db.Exec to receive sql.Result types
package tests

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"taskservice/models"
	r "taskservice/repository"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func loadDB() (*sql.DB, string) {
	err := godotenv.Load()
	if err != nil {
		return nil, "Error loading .env file"
	}

	var db *sql.DB
	cfg := mysql.Config{
		User:   os.Getenv("TEST_DB_USER"),
		Passwd: os.Getenv("TEST_DB_PASSWD"),
		Net:    os.Getenv("TEST_DB_NET"),
		Addr:   os.Getenv("TEST_DB_ADDR"),
		DBName: os.Getenv("TEST_DB_NAME"),
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, "Invalid login to test database"
	}
	return db, ""
}

func TestNewTaskRepo(t *testing.T) {
	prefix := "T(func NewTaskRepo): "
	db, strErr := loadDB()
	if strErr != "" {
		t.Error(prefix, strErr)
	}
	defer db.Close()

	tr := r.NewTaskRepo(db)
	err := tr.DB.Ping()
	if err != nil {
		t.Error(prefix, "Could not connect to database")
	}
}

func TestCreateNew(t *testing.T) {
	prefix := "T(func CreateNew):"
	db, strErr := loadDB()
	if strErr != "" {
		t.Error(prefix, strErr)
	}
	defer db.Close()

	tr := r.NewTaskRepo(db)

	testcases := []models.Task{
		{
			Title:      "T1",
			Priority:   true,
			CreatorID:  100,
			AssigneeID: 1,
			CreatedAt:  time.Now().Format(time.DateTime),
		},
		{
			Title:     "T2",
			Priority:  false,
			CreatorID: 9,
			CreatedAt: time.Now().Format(time.DateTime),
		},
		{
			Title:      "T3",
			Priority:   false,
			CreatorID:  12,
			AssigneeID: 100000,
			CreatedAt:  time.Now().Format(time.DateTime),
		},
	}

	for _, task := range testcases {
		err := tr.CreateNew(&task)
		if err != nil {
			t.Errorf("%s task: %v, err: %v",
				prefix, task, err)
		}
	}
}

func TestGetByID(t *testing.T) {
	prefix := "T(func GetById):"
	db, strErr := loadDB()
	if strErr != "" {
		t.Error(prefix, strErr)
	}
	defer db.Close()

	tr := r.NewTaskRepo(db)

	testcases := map[uint64]error{
		1: nil,
		2: nil,
		3: nil,
		4: nil,
		5: nil,
		0: sql.ErrNoRows,
	}

	for taskId, expected := range testcases {
		_, err := tr.GetByTaskID(taskId)
		if err != expected {
			t.Errorf("%s taskId: %d, err: %v", prefix, taskId, err)
		}
	}
}

func TestGetAll(t *testing.T) {
	prefix := "T(func GetAll):"
	db, strErr := loadDB()
	if strErr != "" {
		t.Error(prefix, strErr)
	}
	defer db.Close()

	tr := r.NewTaskRepo(db)
	testcases := map[uint64]error{
		1: nil,
		2: nil,
		3: sql.ErrNoRows,
		4: nil,
		5: nil,
	}
	lens := []int{0, 2, 1, 0, 1, 1}

	for userId, expected := range testcases {
		tasks, err := tr.GetAllCreated(10000, uint64(userId))
		if err != nil {
			t.Errorf("%s userId: %d, expected: %v", prefix, userId, expected)
		} else if len(tasks) != lens[userId] {
			t.Errorf("%s userId: %d, num tasks expected: %d", prefix, userId, lens[userId])
		}
	}
}

func TestUpdateExisting(t *testing.T) {
	prefix := "T(func UpdateExisting):"
	db, strErr := loadDB()
	if strErr != "" {
		t.Error(prefix, strErr)
	}
	defer db.Close()

	tr := r.NewTaskRepo(db)

	testcases := map[models.Task]error{
		{
			TaskID:    1,
			Title:     "Leave Company",
			Priority:  true,
			CreatorID: 3,
			CreatedAt: time.Now().Format(time.DateTime),
		}: nil,

		{
			TaskID:    2,
			Title:     "Code Harder",
			Priority:  false,
			CreatorID: 1,
			CreatedAt: time.Now().Format(time.DateTime),
		}: nil,
		{
			TaskID: 1000000,
		}: nil,
		{}: nil,
	}

	for task, expected := range testcases {
		err := tr.UpdateExisting(task.TaskID, &task)
		if err != expected {
			t.Errorf("%s taskId: %d, expected: %v, got: %v", prefix, task.TaskID, expected, err)
		}
	}
}

func TestDeleteByID(t *testing.T) {
	prefix := "T(func DeleteByID):"
	db, strErr := loadDB()
	if strErr != "" {
		t.Error(prefix, strErr)
	}
	defer db.Close()

	tr := r.NewTaskRepo(db)

	testcases := []uint64{1, 2, 3, 4, 5, 6, 7, 8} // will work when running full suite or lone test

	for _, taskId := range testcases {
		err := tr.DeleteByTaskID(taskId)
		if err != nil {
			t.Errorf("%s taskId: %d", prefix, taskId)
		}
	}
}
