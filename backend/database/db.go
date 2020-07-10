package database

import (
	"errors"
	"todo-list-study/backend/models"
)

type Database interface {
	AddTask(task models.Task) []models.Task
	ListTasks() []models.Task
	FindTask(uint32) (models.Task, error)
	DeleteTask(taskID uint32) error
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

// DeleteTask purges a task frm array
func (d *DatabaseImpl) DeleteTask(taskID uint32) error {
	var elementIndex *int

	for i, task := range d.Data {
		if task.ID == taskID {
			elementIndex = &i
			break
		}
	}

	if elementIndex == nil {
		return errors.New("Element does not exist")
	}

	switch *elementIndex {
	case 0:
		d.Data = d.Data[1:]
	case len(d.Data) - 1:
		d.Data = d.Data[:*elementIndex]
	default:
		beginning := d.Data[:*elementIndex]
		end := d.Data[*elementIndex+1:]
		d.Data = append(beginning, end...)
	}

	return nil
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
