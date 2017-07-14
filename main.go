package main

import (
	"time"

	"github.com/olioapps/todo-golang/api"
	"github.com/olioapps/todo-golang/service"
)

func main() {
	coreAPI := api.NewCoreAPI()
	todoService := service.NewTodoService(*coreAPI)
	todoService.Start()
	for true {
		time.Sleep(10 * time.Second)
	}
}
