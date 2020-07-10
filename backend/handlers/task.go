package handlers

import (
	"net/http"
	"strconv"
	"time"
	"todo-list-study/backend/database"
	"todo-list-study/backend/models"

	"github.com/gin-gonic/gin"
)

// TaskHandler interface
type TaskHandler interface {
	CreateTask(c *gin.Context)
	ListAllTasks(c *gin.Context)
	FindTask(c *gin.Context)
	DeleteTask(c *gin.Context)
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

	bindErr := c.BindJSON(&task)
	if bindErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid Data",
		})
	}

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

// FindTask locates and return chosen task
func (h *TaskHandlerImpl) FindTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("taskID"), 10, 32)

	if err != nil || c.Param("taskID") == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "No ID given",
		})
	}

	task, err := h.database.FindTask(uint32(taskID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"taskID": taskID,
			"error":  "Could not find task with given id",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

// DeleteTask deletes a single task
func (h *TaskHandlerImpl) DeleteTask(c *gin.Context) {
	taskID, jsonErr := strconv.ParseUint(c.Param("taskID"), 10, 64)

	if jsonErr != nil || c.Param("taskID") == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "No ID given",
		})
	}

	delErr := h.database.DeleteTask(uint32(taskID))
	if delErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"taskID": taskID,
			"error":  "Could not find task with given id",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": models.Task{ID: uint32(taskID)},
	})
}
