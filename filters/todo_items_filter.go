package filters

type TodoItemsFilter struct {
	*BaseTodoFilter
	Name       string
	TodoListID int64
}

func NewTodoItemsFilter() *TodoItemsFilter {
	return &TodoItemsFilter{BaseTodoFilter: &BaseTodoFilter{}}
}
