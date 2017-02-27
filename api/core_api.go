package api

import (
	"github.com/Ligerlilly/todo-golang/dao"
	"github.com/Ligerlilly/todo-golang/db"
	"github.com/Ligerlilly/todo-golang/filters"
	"github.com/Ligerlilly/todo-golang/models"
	log "github.com/Sirupsen/logrus"
	olioAPI "github.com/rachoac/service-skeleton-go/olio/api"
	olioDAO "github.com/rachoac/service-skeleton-go/olio/dao"
)

type TodoItemAPIType interface {
	GetTodoItems(accessContext *models.AccessContext, filter *filters.TodoItemFilter) ([]models.TodoItem, *olioAPI.Exception)
}

type CoreAPI struct {
	olioAPI.OlioBaseCoreAPI
	TodoItemAPI TodoItemAPIType
}

func NewCoreAPI() *CoreAPI {
	log.Info("Initializing todo core api")

	api := CoreAPI{}
	log.Info("Initializing database connection pool")
	connectionManager := olioDAO.NewConnectionManager()
	api.ConnectionManager = connectionManager
	api.TodoItemAPI = NewTodoItemAPI(&api, dao.NewTodoItemDAO(connectionManager))

	migrations := db.NewMigrationsContainer(connectionManager).GetMigrations()

	api.RunMigrations(migrations)

	log.Info("Initializing appplication context")

	return &api
}
