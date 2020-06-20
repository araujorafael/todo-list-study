package handlers_test

import (
	"testing"
	"todo-list/backend/database"
	"todo-list/backend/models"

	"github.com/go-playground/assert/v2"
)

func TestCreateTask(t *testing.T) {
	var db database.Database
	db = new(database.DatabaseImpl)

	task1 := models.Task{
		Title: "test title",
	}

	task2 := models.Task{
		Title: "test title 2",
	}

	db.AddTask(task2)
	resp := db.AddTask(task1)

	assert.Equal(t, 2, len(resp))
}
