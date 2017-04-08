package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ligerlilly/todo-golang/api"
	"github.com/ligerlilly/todo-golang/service/resources"
	"github.com/rachoac/service-skeleton-go/olio/service"
	"github.com/rachoac/service-skeleton-go/olio/service/middleware"
)

type TodoService struct {
	olioService *service.OlioBaseService
}

// type TodoUserExtractor struct {
// 	coreAPI api.CoreAPI
// }

// func NewTodoUserExtractor(coreAPI api.CoreAPI) TodoUserExtractor {
// 	return TodoUserExtractor{coreAPI}
// }

// func (e TodoUserExtractor) ExtractUser(username string, password string, requestId string) (interface{}, error) {
// 	context := api.CreateSystemAccessContext(requestId)

// 	user, err := e.coreAPI.UsersAPI.GetUserByUsername(context, username)
// 	if err != nil || user == nil {
// 		return nil, olioAPI.NewUnauthorizedException("could not find user")
// 	}

// 	return user, nil
// }

// func (e TodoUserExtractor) ExtractUserByUsername(username string, requestId string) (interface{}, error) {
// 	context := api.CreateSystemAccessContext(requestId)

// 	user, exception := e.coreAPI.UsersAPI.GetUserByUsername(context, username)
// 	if exception != nil || user == nil {
// 		return nil, olioAPI.NewUnauthorizedException("Could not find user")
// 	}

// 	return user, nil
// }

// type TodoTokenValidator struct{}

// func (t TodoTokenValidator) IsTokenBlacklisted(token string) (bool, error) {
// 	return false, nil
// }

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
