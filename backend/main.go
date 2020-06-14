package main

import (
	"github.com/gin-gonic/gin"
)

type PayloadTest struct {
	Title      string `json:"title"`
	Message    string `json:"message"`
	BudegaMode bool   `json:"budega_mode"`
}

func pingHandler(c *gin.Context) {
	la := PayloadTest{
		Title:      "Some Title",
		Message:    "Some Message",
		BudegaMode: false,
	}

	c.JSON(200, la)
}

func router(r *gin.Engine) *gin.Engine {
	r.GET("/ping", pingHandler)
	return r
}

func main() {
	server := router(gin.Default())
	server.Run("0.0.0.0:3000")
}
