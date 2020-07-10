package router

import (
	"todo-list-study/backend/handlers"

	"github.com/gin-gonic/gin"
)

// BuildRouter create all routes as needed
func BuildRouter(server *gin.Engine,
	taskHandler handlers.TaskHandler,
	helloExampleHandler handlers.HelloExampleHandler) *gin.Engine {

	server.GET("/ping", helloExampleHandler.Ping)

	tasksRouter := server.Group("/tasks")
	{
		tasksRouter.POST("", taskHandler.CreateTask)
		tasksRouter.GET("", taskHandler.ListAllTasks)
		tasksRouter.GET("/:taskID", taskHandler.FindTask)
		tasksRouter.DELETE("/:taskID", taskHandler.DeleteTask)
	}

	return server
}
