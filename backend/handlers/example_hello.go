package handlers

import (
	"github.com/gin-gonic/gin"
)

// HelloExampleHandler interface definition
type HelloExampleHandler interface {
	Ping(c *gin.Context)
}

// payloadTest default struct that will be used to send
type payloadTest struct {
	Title      string `json:"title"`
	Message    string `json:"message"`
	BudegaMode bool   `json:"budega_mode"`
}

// HelloExampleHandler responsable for interface method implementation
type HelloExampleHandlerImpl struct{}

// Ping handler function :)
func (h HelloExampleHandlerImpl) Ping(c *gin.Context) {
	hello := payloadTest{
		Title:      "Some Title",
		Message:    "Some Message",
		BudegaMode: false,
	}

	c.JSON(200, hello)
}
