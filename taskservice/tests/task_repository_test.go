package tests

import (
	"database/sql"
	"testing"
	"time"

	"taskservice/models"
	r "taskservice/repository"
)

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
		ctx, cancel := loadContext()
		defer cancel()

		err := tr.CreateNew(ctx, &task)
		if err != nil {
			t.Errorf("%s task: %v, err: %v", prefix, task, err)
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
		ctx, cancel := loadContext()
		defer cancel()

		_, err := tr.GetByTaskID(ctx, taskId)
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
		ctx, cancel := loadContext()
		defer cancel()

		tasks, err := tr.GetAllCreated(ctx, 10000, uint64(userId))
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
		ctx, cancel := loadContext()
		defer cancel()

		err := tr.UpdateExisting(ctx, task.TaskID, &task)
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
		ctx, cancel := loadContext()
		defer cancel()

		err := tr.DeleteByTaskID(ctx, taskId)
		if err != nil {
			t.Errorf("%s taskId: %d", prefix, taskId)
		}
	}
}
