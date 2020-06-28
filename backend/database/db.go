package database

import (
	"todo-list-study/backend/models"
)

type Database interface {
	AddTask(task models.Task) []models.Task
	ListTasks() []models.Task
}

type DatabaseImpl struct {
	Data []models.Task
}

func NewDatabase() *DatabaseImpl {
	return &DatabaseImpl{}
}

func (d *DatabaseImpl) AddTask(task models.Task) []models.Task {
	d.Data = append(d.Data, task)
	return d.Data
}

func (d *DatabaseImpl) ListTasks() []models.Task {
	return d.Data
}
