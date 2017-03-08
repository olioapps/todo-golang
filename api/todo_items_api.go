package api

import (
	"github.com/Ligerlilly/todo-golang/dao"
	"github.com/Ligerlilly/todo-golang/filters"
	"github.com/Ligerlilly/todo-golang/models"
	olioAPI "github.com/rachoac/service-skeleton-go/olio/api"
	"github.com/siddontang/go/log"
)

type TodoItemsAPI struct {
	applicationContext *CoreAPI
	dao                *dao.TodoItemsDAO
}

func NewTodoItemsAPI(coreAPI *CoreAPI, dao *dao.TodoItemsDAO) *TodoItemsAPI {
	todoItemsAPI := TodoItemsAPI{}
	todoItemsAPI.applicationContext = coreAPI
	todoItemsAPI.dao = dao
	return &todoItemsAPI
}

func (tia *TodoItemsAPI) GetTodoItems(filter *filters.TodoItemsFilter) ([]models.TodoItem, *olioAPI.Exception) {
	todos, err := tia.dao.GetTodoItems(filter)

	if err != nil {
		log.Error(err)
		return nil, olioAPI.NewRuntimeException(err.Error())
	}

	return todos, nil
}
