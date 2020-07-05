package database

import (
	"errors"
	"todo-list-study/backend/models"
)

type Database interface {
	AddTask(task models.Task) []models.Task
	ListTasks() []models.Task
	FindTask(uint32) (models.Task, error)
}

// DatabaseImpl Databse "tables" implementation
type DatabaseImpl struct {
	Data []models.Task
}

// NewDatabase instanciates a new database
func NewDatabase() *DatabaseImpl {
	return &DatabaseImpl{}
}

// AddTask adds a task into database
func (d *DatabaseImpl) AddTask(task models.Task) []models.Task {
	task.ID = uint32(len(d.Data))
	d.Data = append(d.Data, task)
	return d.Data
}

// ListTasks return all tasks saved on database
func (d *DatabaseImpl) ListTasks() []models.Task {
	return d.Data
}

func (d *DatabaseImpl) FindTask(id uint32) (models.Task, error) {
	task, err := d.taskListFilter(d.Data, id)
	return task, err
}

func (d *DatabaseImpl) taskListFilter(tasks []models.Task, id uint32) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return models.Task{}, errors.New("Task Not Found")
}
