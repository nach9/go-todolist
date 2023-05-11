package dtoTodo

import entityTodo "github.com/nach9/go-todolist/pkg/todos/entity"

type TodoResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    entityTodo.Todo `json:"data"`
}

type TodoListResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    []entityTodo.Todo `json:"data"`
}
