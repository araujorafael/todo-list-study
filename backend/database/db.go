package database

import (
	"todo-list/backend/models"
)

type Database interface {
	AddTask(task models.Task) []models.Task
	ListTasks() []models.Task
}

type DatabaseImpl struct {
	Data []models.Task
}

func (d *DatabaseImpl) AddTask(task models.Task) []models.Task {
	d.Data = append(d.Data, task)
	return d.Data
}

func (d *DatabaseImpl) ListTasks() []models.Task {
	return d.Data
}
