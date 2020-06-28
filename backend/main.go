package main

import (
	"fmt"

	"todo-list-study/backend/database"
	"todo-list-study/backend/handlers"
	"todo-list-study/backend/router"

	"github.com/gin-gonic/gin"
)

// main it's app entry point and orchestrate dependencies
func main() {
	var serverConfs = gin.Default()
	var systemCommonDB = database.NewDatabase()

	// instanciate handlers
	taskHandler := handlers.BuildTaskHandler(systemCommonDB)
	helloExample := new(handlers.HelloExampleHandlerImpl)

	server := router.BuildRouter(serverConfs, taskHandler, helloExample)
	err := server.Run("0.0.0.0:3000")
	if err != nil {
		fmt.Println("error at strting server\n", err)
	}
}
