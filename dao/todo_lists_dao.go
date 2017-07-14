package dao

import (
	log "github.com/Sirupsen/logrus"
	"github.com/olioapps/todo-golang/filters"
	"github.com/olioapps/todo-golang/models"
	olioDAO "github.com/olioapps/service-skeleton-go/olio/dao"
)

type TodoListsDAO struct {
	BaseDAO
}

func NewTodoListsDAO(connectionManager olioDAO.ConnectionProvider) *TodoListsDAO {
	dao := TodoListsDAO{
		BaseDAO{connectionManager},
	}

	return &dao
}

func (tld *TodoListsDAO) GetTodoLists(filter *filters.TodoListsFilter) ([]models.TodoList, error) {
	db := tld.connectionManager.GetDb()

	var filterConditions map[string]interface{} = make(map[string]interface{})

	var todoLists []models.TodoList
	db = db.Where(filterConditions)
	db = db.Find(&todoLists)

	todoItemsDAO := NewTodoItemsDAO(tld.connectionManager)

	for i, list := range todoLists {
		list.TodoItems = make([]models.TodoItem, 0)
		todoItemFilter := filters.NewTodoItemsFilter()
		todoItemFilter.TodoListID = list.ID
		todoItems, err := todoItemsDAO.GetTodoItems(todoItemFilter)
		if err != nil {
			log.Error("Failed to get todoItems for list ID", list.ID, err)
		}

		todoLists[i].TodoItems = todoItems
	}

	return todoLists, db.Error
}
