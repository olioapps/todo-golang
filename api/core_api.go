package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/olioapps/todo-golang/dao"
	"github.com/olioapps/todo-golang/db"
	"github.com/olioapps/todo-golang/filters"
	"github.com/olioapps/todo-golang/models"
	olioAPI "github.com/olioapps/service-skeleton-go/olio/api"
	olioDAO "github.com/olioapps/service-skeleton-go/olio/dao"
	olioUtil "github.com/olioapps/service-skeleton-go/olio/util"
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

	connectionString := olioUtil.GetEnv("DB_CONNECTION_STRING", "")
	if connectionString == "" {
		panic("No DB_CONNECTION_STRING (eg. username:password@/databasename?parseTime=true) provided; cannot start service")
	}

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
