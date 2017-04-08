package dao

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ligerlilly/todo-golang/filters"
	"github.com/ligerlilly/todo-golang/models"
	olioDAO "github.com/rachoac/service-skeleton-go/olio/dao"
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
		log.Infof("=================  %v", todoItems)
		if err != nil {
			log.Error("Failed to get todoItems for list ID", list.ID, err)
		}

		todoLists[i].TodoItems = todoItems
	}

	log.Infof("todoLists=================  %+v", todoLists)

	return todoLists, db.Error
}
