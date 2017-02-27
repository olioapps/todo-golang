package filters

type TodoItemFilter struct {
	*BaseTodoFilter
	name string
}

func NewTodoItemFilter() *TodoItemFilter {
	return &TodoItemFilter{BaseTodoFilter: &BaseTodoFilter{}}
}
