package main

import (
	"todo-list/backend/handlers"
	"todo-list/backend/router"

	"github.com/gin-gonic/gin"
)

// main it's app entry point and orchestrate dependencies
func main() {
	var serverConfs = gin.Default()

	// instanciate handlers
	var taskHandler = new(handlers.TaskHandlerImpl)
	var helloExample = new(handlers.HelloExampleHandlerImpl)

	server := router.BuildRouter(serverConfs, taskHandler, helloExample)
	server.Run("0.0.0.0:3000")
}
