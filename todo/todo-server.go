package todo

import (
	"log"
	"net"

	"golang.org/x/net/context"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/ligerlilly/todo-golang/api"
	"github.com/ligerlilly/todo-golang/filters"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

func (s *server) AddTodoItem(context.Context, *TodoItem) (*google_protobuf.Empty, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *server) ListTodos(context.Context, *google_protobuf.Empty) (*TodoList, error) {
	coreAPI := api.NewCoreAPI()
	filter := filters.NewTodoListsFilter()
	todoLists, err := coreAPI.TodoListsAPI.GetTodoLists(filter)
	if err != nil {
		return nil, err
	}

	list := &TodoList{Name: todoLists[0].Name}
	return list, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterDoSomethingServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
