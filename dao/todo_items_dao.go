package dao

import (
	"github.com/olioapps/todo-golang/filters"
	"github.com/olioapps/todo-golang/models"
	olioDAO "github.com/olioapps/service-skeleton-go/olio/dao"
)

type TodoItemsDAO struct {
	BaseDAO
}

func NewTodoItemsDAO(connectionManager olioDAO.ConnectionProvider) *TodoItemsDAO {
	dao := TodoItemsDAO{
		BaseDAO{connectionManager},
	}

	return &dao
}

func (td *TodoItemsDAO) GetTodoItems(filter *filters.TodoItemsFilter) ([]models.TodoItem, error) {
	db := td.connectionManager.GetDb()

	var filterConditions map[string]interface{} = make(map[string]interface{})

	if filter.TodoListID > 1 {
		filterConditions["todo_list_id"] = filter.TodoListID
	}

	var results []models.TodoItem
	db = db.Where(filterConditions)
	db = db.Find(&results)

	return results, db.Error
}
