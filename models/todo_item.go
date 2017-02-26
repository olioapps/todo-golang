package models

type TodoItem struct {
	BaseModel
	Name string `json:"name,omitempty"`
}
