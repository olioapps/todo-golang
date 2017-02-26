package api

import (
	log "github.com/Sirupsen/logrus"
	olioAPI "github.com/rachoac/service-skeleton-go/olio/api"
	olioDAO "github.com/rachoac/service-skeleton-go/olio/dao"
	"github.com/sim-works/sim-backend/db"
	"github.com/thedataguild/faer/models"
)

type TodoAPIType interface {
	GetTodos(accessContext *models.AccessContext, filter *filter.TodosFilter)
}

type CoreAPI struct {
	olioAPI.OlioBaseCoreAPI
	TodoAPI TodoAPIType
}

func NewCoreAPI() *CoreAPI {
	log.Info("Initializing todo core api")

	api := CoreAPI{}
	log.Info("Initializing database connection pool")
	connectionManager := olioDAO.NewConnectionManager()
	api.ConnectionManager = connectionManager
	api.TodoAPI = NewTodoAPI(&api, dao.NewTodoDAO(connectionManager))

	migrations := db.NewMigrationsContainer(connectionManager).GetMigrations()

	api.RunMigrations(migrations)

	log.Info("Initializing appplication context")

	return &api
}
