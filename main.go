package main

import (
	"time"

	"github.com/ligerlilly/todo-golang/api"
	"github.com/ligerlilly/todo-golang/service"
)

func main() {
	coreAPI := api.NewCoreAPI()
	todoService := service.NewTodoService(*coreAPI)
	todoService.Start()
	for true {
		time.Sleep(10 * time.Second)
	}
}
