package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskHandler interface {
	CreateTask(c *gin.Context)
}

// Task payload definition
type Task struct {
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// responseBody define default response body
type responseBody struct {
	status string
	data   interface{} // FIXME: AVOID USE interface{}
}

// TaskHandlerImpl responsable for interface methods implementation
type TaskHandlerImpl struct{}

// CreateTask responsible for create task
func (h TaskHandlerImpl) CreateTask(c *gin.Context) {
	var task Task

	c.BindJSON(&task)
	createdAt := time.Now()
	task.CreatedAt = &createdAt

	c.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"data":   task,
	})

}
