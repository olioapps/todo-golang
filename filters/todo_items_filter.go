package filters

type TodoItemsFilter struct {
	*BaseTodoFilter
	Name string
}

func NewTodoItemsFilter() *TodoItemsFilter {
	return &TodoItemsFilter{BaseTodoFilter: &BaseTodoFilter{}}
}
