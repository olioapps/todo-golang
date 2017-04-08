package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ligerlilly/todo-golang/dao"
	"github.com/ligerlilly/todo-golang/db"
	"github.com/ligerlilly/todo-golang/filters"
	"github.com/ligerlilly/todo-golang/models"
	olioAPI "github.com/rachoac/service-skeleton-go/olio/api"
	olioDAO "github.com/rachoac/service-skeleton-go/olio/dao"
)

type TodoItemsAPIType interface {
	GetTodoItems(filter *filters.TodoItemsFilter) ([]models.TodoItem, *olioAPI.Exception)
}

type TodoListsAPIType interface {
	GetTodoLists(filter *filters.TodoListsFilter) ([]models.TodoList, *olioAPI.Exception)
}

type CoreAPI struct {
	olioAPI.OlioBaseCoreAPI
	TodoItemsAPI TodoItemsAPIType
	TodoListsAPI TodoListsAPIType
}

func NewCoreAPI() *CoreAPI {
	log.Info("Initializing todo core api")

	api := CoreAPI{}
	log.Info("Initializing database connection pool")
	connectionManager := olioDAO.NewConnectionManager()
	api.ConnectionManager = connectionManager
	api.TodoItemsAPI = NewTodoItemsAPI(&api, dao.NewTodoItemsDAO(connectionManager))
	api.TodoListsAPI = NewTodoListsAPI(&api, dao.NewTodoListsDAO(connectionManager))

	migrations := db.NewMigrationsContainer(connectionManager).GetMigrations()

	api.RunMigrations(migrations)

	log.Info("Initializing appplication context")

	return &api
}
