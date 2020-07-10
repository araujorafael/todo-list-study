package handlers_test

import (
	"testing"

	"todo-list-study/backend/database"
	"todo-list-study/backend/models"

	"github.com/go-playground/assert/v2"
)

func TestCreateTask(t *testing.T) {
	db := database.DatabaseImpl{}

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

func TestFindTask(t *testing.T) {
	db := database.DatabaseImpl{
		Data: []models.Task{
			{
				ID:    1,
				Title: "test title",
			},
			{
				ID:    2,
				Title: "test title 2",
			},
		},
	}

	resp, err := db.FindTask(2)

	assert.Equal(t, uint32(2), resp.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteTask(t *testing.T) {
	t.Run("remove middle item", func(t *testing.T) {
		db := database.DatabaseImpl{
			Data: []models.Task{
				{ID: 1, Title: "test title"},
				{ID: 2, Title: "test title 2"},
				{ID: 3, Title: "test title 3"},
			},
		}

		err := db.DeleteTask(2)

		assert.Equal(t, nil, err)
		assert.Equal(t, uint32(1), db.Data[0].ID)
		assert.Equal(t, 2, len(db.Data))
	})

	t.Run("Should remove first item", func(t *testing.T) {
		db := database.DatabaseImpl{
			Data: []models.Task{
				{ID: 1, Title: "test title"},
				{ID: 2, Title: "test title 2"},
				{ID: 3, Title: "test title 3"},
			},
		}

		err := db.DeleteTask(1)

		assert.Equal(t, nil, err)
		assert.Equal(t, uint32(2), db.Data[0].ID)
		assert.Equal(t, 2, len(db.Data))
	})

	t.Run("Should remove middle item", func(t *testing.T) {
		db := database.DatabaseImpl{
			Data: []models.Task{
				{ID: 1, Title: "test title"},
				{ID: 2, Title: "test title 2"},
				{ID: 3, Title: "test title 3"},
			},
		}

		err := db.DeleteTask(3)

		assert.Equal(t, nil, err)
		assert.Equal(t, uint32(2), db.Data[len(db.Data)-1].ID)
		assert.Equal(t, 2, len(db.Data))
	})

	t.Run("Should return error when requested elemente does not exist",
		func(t *testing.T) {
			db := database.DatabaseImpl{
				Data: []models.Task{
					{ID: 1, Title: "test title"},
					{ID: 2, Title: "test title 2"},
					{ID: 3, Title: "test title 3"},
				},
			}

			err := db.DeleteTask(666)

			assert.NotEqual(t, nil, err)
			assert.Equal(t, 3, len(db.Data))
		})
}
