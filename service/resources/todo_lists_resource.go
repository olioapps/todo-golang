package resources

import (
	"context"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/ligerlilly/todo-golang/api"
	"github.com/ligerlilly/todo-golang/filters"
	todo "github.com/ligerlilly/todo-golang/todo"
	olioMiddleware "github.com/rachoac/service-skeleton-go/olio/service/middleware"
	"google.golang.org/grpc"
)

type TodoListsResource struct {
	BaseTodoResource
	coreAPI    *api.CoreAPI
	Connection grpc.ClientConn
}

func NewTodoListsResource(coreAPI *api.CoreAPI) *TodoListsResource {
	obj := TodoListsResource{}
	obj.coreAPI = coreAPI
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", "localhost", "50051"), grpc.WithInsecure())
	if err != nil {
		return nil
	}

	obj.Connection = *conn

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

	// todoLists, exception := resource.coreAPI.TodoListsAPI.GetTodoLists(filter)
	// if exception != nil {
	// 	resource.ReturnError(c, exception.ErrorCode, exception.Err)
	// 	return
	// }

	client := todo.NewDoSomethingClient(&resource.Connection)
	result, err := client.ListTodos(context.Background(), &google_protobuf.Empty{})
	if err != nil {
		resource.ReturnError(c, 500, err.Error())
	}

	resource.ReturnJSON(c, 200, result)
}
