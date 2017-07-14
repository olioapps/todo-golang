package api

import (
	"github.com/olioapps/todo-golang/dao"
	"github.com/olioapps/todo-golang/filters"
	"github.com/olioapps/todo-golang/models"
	log "github.com/Sirupsen/logrus"
	olioAPI "github.com/olioapps/service-skeleton-go/olio/api"
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
