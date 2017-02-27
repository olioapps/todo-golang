package api

import (
	"github.com/Ligerlilly/todo-golang/dao"
	"github.com/Ligerlilly/todo-golang/filters"
	"github.com/Ligerlilly/todo-golang/models"
	olioAPI "github.com/rachoac/service-skeleton-go/olio/api"
	"github.com/siddontang/go/log"
)

type TodoItemAPI struct {
	applicationContext *CoreAPI
	dao                *dao.TodoItemDAO
}

func NewTodoItemAPI(coreAPI *CoreAPI, dao *dao.TodoItemDAO) *TodoItemAPI {
	todoItemAPI := TodoItemAPI{}
	todoItemAPI.applicationContext = coreAPI
	todoItemAPI.dao = dao
	return &todoItemAPI
}

func (ta *TodoItemAPI) GetTodoItems(accessContext *models.AccessContext, filter *filters.TodoItemFilter) ([]models.TodoItem, *olioAPI.Exception) {
	todos, err := ta.dao.GetTodoItems(filter)

	if err != nil {
		log.Error(err)
		return nil, olioAPI.NewRuntimeException(err.Error())
	}

	return todos, nil
}
