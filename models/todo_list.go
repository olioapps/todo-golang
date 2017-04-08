package models

type TodoList struct {
	BaseModel
	Name      string     `json:"name,omitempty"`
	TodoItems []TodoItem `json:"todoItems,omitempty"`
}
