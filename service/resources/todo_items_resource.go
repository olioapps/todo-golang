package resources

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/olioapps/todo-golang/api"
	"github.com/olioapps/todo-golang/filters"
	olioMiddleware "github.com/olioapps/service-skeleton-go/olio/service/middleware"
)

type TodoItemsResource struct {
	BaseTodoResource
	coreAPI *api.CoreAPI
}

func NewTodoItemsResource(coreAPI *api.CoreAPI) *TodoItemsResource {
	obj := TodoItemsResource{}
	obj.coreAPI = coreAPI
	return &obj
}

func (resource TodoItemsResource) Init(e *gin.Engine, whiteList *olioMiddleware.WhiteList) {
	log.Debug("Setting up todo items resource")

	e.GET("/api/v1/todo_items", resource.getTodoItems)
}

func (resource TodoItemsResource) getTodoItems(c *gin.Context) {
	filter := filters.NewTodoItemsFilter()
	err := resource.parseFilter(c, filter.BaseTodoFilter)
	if err != nil {
		resource.ReturnError(c, 400, err.Error())
	}

	if name := resource.ParseString(c, "name"); name != "" {
		filter.Name = name
	}

	todoItems, exception := resource.coreAPI.TodoItemsAPI.GetTodoItems(filter)
	if exception != nil {
		resource.ReturnError(c, exception.ErrorCode, exception.Err)
		return
	}

	resource.ReturnJSON(c, 200, todoItems)

}
