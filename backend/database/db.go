package database

import (
	"todo-list/backend/models"
)

type Database interface {
	AddTask(task models.Task) []models.Task
}

type DatabaseImpl struct {
	Data []models.Task
}

func (d *DatabaseImpl) AddTask(task models.Task) []models.Task {
	d.Data = append(d.Data, task)
	return d.Data
}
