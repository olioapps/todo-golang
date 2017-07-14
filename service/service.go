package service

import (
	"github.com/gin-gonic/gin"
	"github.com/olioapps/todo-golang/api"
	"github.com/olioapps/todo-golang/service/resources"
	"github.com/olioapps/service-skeleton-go/olio/service"
	"github.com/olioapps/service-skeleton-go/olio/service/middleware"
)

type TodoService struct {
	olioService *service.OlioBaseService
}

func NewTodoService(coreAPI api.CoreAPI) TodoService {
	olioService := service.New()
	todoService := TodoService{}
	todoService.olioService = olioService

	corsMiddleware := middleware.NewOlioCORSMiddleware().Create()

	middlewares := []gin.HandlerFunc{
		corsMiddleware,
	}

	whitelist := middleware.NewWhitelist()
	olioService.Init(
		whitelist,

		// middleware
		middlewares,

		// resources
		[]service.OlioResourceHandler{
			resources.NewTodoItemsResource(&coreAPI),
			resources.NewTodoListsResource(&coreAPI),
		})

	return todoService
}

func (service TodoService) Start() {
	service.olioService.Start()
}

func (service TodoService) Stop() {
	service.olioService.Stop()
}
