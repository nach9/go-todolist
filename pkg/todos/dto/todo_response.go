package dtoTodo

import entityTodo "github.com/nach9/go-todolist/pkg/todos/entity"

type TodoResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    entityTodo.Todo `json:"data,omitempty"`
}

type TodoListResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    []entityTodo.Todo `json:"data"`
}

type TodoDeleteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    *Blank `json:"data"`
}

type Blank struct {
	ID int64 `json:"id,omitempty"`
}
