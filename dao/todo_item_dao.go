package dao

import (
	olioDAO "github.com/rachoac/service-skeleton-go/olio/dao"
)

type TodoItemDAO struct {
	BaseDAO
}

func NewTodoItemDAO(connectionManager olioDAO.ConnectionProvider) *TodoItemDAO {
	dao := TodoItemDAO{
		BaseDAO{connectionManager},
	}
	return &dao
}

func (td *TodoItemDAO) GetTodoItems(filter *filter.TodoItemFilter) ([]models.Todo, error) {
	db := td.connectionManager.GetDb()

	var filterConditions map[string]interface{} = make(map[string]interface{})

	var results []models.TodoItem
	db = db.Where(filterConditions)

	return results, db.Error
}
