package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"todo-list-study/backend/database"
	"todo-list-study/backend/handlers"
	"todo-list-study/backend/models"
	"todo-list-study/backend/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

type responseJson struct {
	Status string      `json:"status"`
	Data   models.Task `json:"data"`
}

type responseListJson struct {
	Data []models.Task `json:"data"`
}

// -------- test helpers
type CustomInternalTestError struct {
	Err  error
	Data interface{}
}

func (e *CustomInternalTestError) Error() string {
	errorStr := `An error occurred at assertions:
Error: %s
Input: %+v`
	return fmt.Sprintf(errorStr, e.Err.Error(), e.Data)
}

func NewCustomTestError(data interface{}, err error) CustomInternalTestError {
	return CustomInternalTestError{
		Err:  err,
		Data: data,
	}
}

//-----------------
func TestCreateTask(t *testing.T) {
	var serverConfs = gin.Default()

	db := new(database.DatabaseImpl)
	var taskHandler = handlers.BuildTaskHandler(db)
	var helloExample = new(handlers.HelloExampleHandlerImpl)

	router := router.BuildRouter(serverConfs, taskHandler, helloExample)

	payload := strings.NewReader(`{"message": "test", "title": "title test"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/task", payload)
	router.ServeHTTP(w, req)

	resp := responseJson{}
	jsonErr := json.Unmarshal(w.Body.Bytes(), &resp)
	if jsonErr != nil {
		err := NewCustomTestError(w.Body.String(), jsonErr)

		t.Errorf("Could not parse body response\n %s", err.Error())
	}

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, resp.Status, "created")
	assert.Equal(t, resp.Data.Message, "test")
	assert.Equal(t, resp.Data.Title, "title test")
	assert.NotEqual(t, resp.Data.CreatedAt, nil)
}

func TestListAllTasks(t *testing.T) {
	var serverConfs = gin.Default()

	dbMock := []models.Task{
		{Title: "Title task 1"},
		{Title: "Title task 2"},
	}
	db := database.DatabaseImpl{
		Data: dbMock,
	}

	var taskHandler = handlers.BuildTaskHandler(&db)
	var helloExample = new(handlers.HelloExampleHandlerImpl)

	router := router.BuildRouter(serverConfs, taskHandler, helloExample)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	var resp responseListJson
	errJson := json.Unmarshal(w.Body.Bytes(), &resp)
	if errJson != nil {
		err := NewCustomTestError(w.Body.String(), errJson)
		t.Errorf("Could not parse body response\n %s", err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(resp.Data))
	assert.Equal(t, dbMock, resp.Data)
}
