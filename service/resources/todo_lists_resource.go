package resources

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/ligerlilly/todo-golang/api"
	"github.com/ligerlilly/todo-golang/filters"
	olioMiddleware "github.com/rachoac/service-skeleton-go/olio/service/middleware"
)

type TodoListsResource struct {
	BaseTodoResource
	coreAPI *api.CoreAPI
}

func NewTodoListsResource(coreAPI *api.CoreAPI) *TodoListsResource {
	obj := TodoListsResource{}
	obj.coreAPI = coreAPI
	return &obj
}

func (resource TodoListsResource) Init(e *gin.Engine, whiteList *olioMiddleware.WhiteList) {
	log.Debug("Setting up todo lists resource")

	e.GET("/api/v1/todo_lists", resource.getTodoLists)
}

func (resource TodoListsResource) getTodoLists(c *gin.Context) {
	filter := filters.NewTodoListsFilter()
	err := resource.parseFilter(c, filter.BaseTodoFilter)
	if err != nil {
		resource.ReturnError(c, 400, err.Error())
	}

	if name := resource.ParseString(c, "name"); name != "" {
		filter.Name = name
	}

	todoLists, exception := resource.coreAPI.TodoListsAPI.GetTodoLists(filter)
	if exception != nil {
		resource.ReturnError(c, exception.ErrorCode, exception.Err)
		return
	}

	resource.ReturnJSON(c, 200, todoLists)
}
