package main

import (
	"time"

	"github.com/ligerlilly/todo-golang/api"
	"github.com/ligerlilly/todo-golang/service"
	"github.com/ligerlilly/todo-golang/todo"
)

func main() {
	coreAPI := api.NewCoreAPI()
	todoService := service.NewTodoService(*coreAPI)
	todoService.Start()
	todo.StartServer()

	for true {
		time.Sleep(10 * time.Second)
	}
}
