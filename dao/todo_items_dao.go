package dao

import (
	"github.com/Ligerlilly/todo-golang/filters"
	"github.com/Ligerlilly/todo-golang/models"
	olioDAO "github.com/rachoac/service-skeleton-go/olio/dao"
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

	var results []models.TodoItem
	db = db.Where(filterConditions)
	db = db.Find(&results)

	return results, db.Error
}
