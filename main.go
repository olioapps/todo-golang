package main

import (
	"time"

	"github.com/Ligerlilly/todo-golang/api"
)

func main() {
	coreAPI := api.NewCoreAPI()
	todoService := service.NewTodoService(*coreAPI)
	todoService.start()
	for true {
		time.Sleep(10 * time.Second)
	}
}
