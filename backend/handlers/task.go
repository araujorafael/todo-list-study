package handlers

import (
	"net/http"
	"time"
	"todo-list/backend/database"
	"todo-list/backend/models"

	"github.com/gin-gonic/gin"
)

// TaskHandler interface
type TaskHandler interface {
	CreateTask(c *gin.Context)
	ListAllTasks(c *gin.Context)
}

// responseBody define default response body
type responseBody struct {
	status string
	data   interface{} // FIXME: AVOID USE interface{}
}

// TaskHandlerImpl responsable for interface methods implementation
type TaskHandlerImpl struct {
	database database.Database
}

// BuildTaskHandler Create and inject all dependencies needed to
// build a instance of TaskHandler implementation
func BuildTaskHandler(db database.Database) *TaskHandlerImpl {
	return &TaskHandlerImpl{
		database: db,
	}
}

// CreateTask responsible for create task
func (h *TaskHandlerImpl) CreateTask(c *gin.Context) {
	var task models.Task

	c.BindJSON(&task)
	createdAt := time.Now()
	task.CreatedAt = &createdAt

	h.database.AddTask(task)

	c.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"data":   task,
	})
}

// ListAllTasks return all tasks saved on DB
func (h *TaskHandlerImpl) ListAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": h.database.ListTasks(),
	})
}
