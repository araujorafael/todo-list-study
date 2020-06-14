package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-list/backend/handlers"
	"todo-list/backend/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

type responseJson struct {
	Status string        `json:"status"`
	Data   handlers.Task `json:"data"`
}

func jsonToStruct(jsonStr string) responseJson {
	var resp responseJson
	json.Unmarshal([]byte(jsonStr), &resp)
	return resp
}

func TestCreateTask(t *testing.T) {
	var serverConfs = gin.Default()
	var taskHandler = new(handlers.TaskHandlerImpl)
	var helloExample = new(handlers.HelloExampleHandlerImpl)
	router := router.BuildRouter(serverConfs, taskHandler, helloExample)

	payload := strings.NewReader(`{"message": "tes", "title": "title test"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/task", payload)
	router.ServeHTTP(w, req)

	resp := jsonToStruct(w.Body.String())
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, resp.Status, "created")
	assert.Equal(t, resp.Data.Message, "tes")
	assert.Equal(t, resp.Data.Title, "title test")
	assert.NotEqual(t, resp.Data.CreatedAt, nil)
}
