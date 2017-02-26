package api

import "github.com/siddontang/go/log"

type TodoAPI struct {
	applicationContext *CoreAPI
	dao                *dao.TodoItemDAO
}

func NewTodoItemAPI(coreAPI *CoreAPI, dao *dao.TodoDAO) *TodoItemAPI {
	todoItemAPI := TodoItemAPI{}
	todoItemAPI.applicationContext = coreAPI
	todoItemAPI.dao = dao
	return &todoAPI
}

func (ta *TodoAPI) GetTodoItems(accessContext *models.AccessContext, filter *filter.TodoFilter) ([]models.Todo, *olioAPI.Exception) {
	todos, err := ta.dao.GetTodoItems(filter)

	if err != nil {
		log.Error(err)
		return nil, olioAPI.NewRuntimeException(err.Error())
	}

	return todos, nil
}
