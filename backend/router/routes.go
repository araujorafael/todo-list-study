package router

import (
	"todo-list/backend/handlers"

	"github.com/gin-gonic/gin"
)

// BuildRouter create all routes as needed
func BuildRouter(server *gin.Engine,
	taskHandler handlers.TaskHandler,
	helloExampleHandler handlers.HelloExampleHandler) *gin.Engine {

	server.GET("/ping", helloExampleHandler.Ping)
	server.POST("/task", taskHandler.CreateTask)

	return server
}
