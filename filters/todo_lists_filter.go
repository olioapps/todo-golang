package filters

type TodoListsFilter struct {
	*BaseTodoFilter
	Name string
}

func NewTodoListsFilter() *TodoListsFilter {
	return &TodoListsFilter{BaseTodoFilter: &BaseTodoFilter{}}
}
