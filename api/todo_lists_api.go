package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/olioapps/todo-golang/dao"
	"github.com/olioapps/todo-golang/filters"
	"github.com/olioapps/todo-golang/models"
	olioAPI "github.com/olioapps/service-skeleton-go/olio/api"
)

type TodoListsAPI struct {
	applicationContext *CoreAPI
	dao                *dao.TodoListsDAO
}

func NewTodoListsAPI(coreAPI *CoreAPI, dao *dao.TodoListsDAO) *TodoListsAPI {
	todoListsAPI := TodoListsAPI{}
	todoListsAPI.applicationContext = coreAPI
	todoListsAPI.dao = dao
	return &todoListsAPI
}

func (tla *TodoListsAPI) GetTodoLists(filter *filters.TodoListsFilter) ([]models.TodoList, *olioAPI.Exception) {
	todoLists, err := tla.dao.GetTodoLists(filter)
	if err != nil {
		log.Error(err)
		return nil, olioAPI.NewRuntimeException(err.Error())
	}

	return todoLists, nil
}
