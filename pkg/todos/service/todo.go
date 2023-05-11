package serviceTodo

import (
	"errors"

	repositoryActivity "github.com/nach9/go-todolist/pkg/activities/repository"
	dtoTodo "github.com/nach9/go-todolist/pkg/todos/dto"
	entityTodo "github.com/nach9/go-todolist/pkg/todos/entity"
	repositoryTodo "github.com/nach9/go-todolist/pkg/todos/repository"
)

type TodoService interface {
	GetAll(param dtoTodo.GetAllParam) []entityTodo.Todo
	GetById(id int64) (entityTodo.Todo, error)
	Create(body dtoTodo.CreateTodoBody) (entityTodo.Todo, error)
	UpdateById(id int64, body dtoTodo.UpdateTodoBody) (entityTodo.Todo, error)
	DeleteById(id int64) (entityTodo.Todo, error)
}

type todoService struct {
	repo    repositoryTodo.TodoRepo
	repoAct repositoryActivity.ActivityRepo
}

func NewTodoService(repo repositoryTodo.TodoRepo, repoAct repositoryActivity.ActivityRepo) TodoService {
	return &todoService{repo, repoAct}
}

func (s todoService) GetAll(param dtoTodo.GetAllParam) []entityTodo.Todo {
	return s.repo.FindAll(param)
}

func (s todoService) GetById(id int64) (entityTodo.Todo, error) {
	return s.repo.FindById(id)
}

func (s todoService) Create(body dtoTodo.CreateTodoBody) (entityTodo.Todo, error) {
	_, err := s.repoAct.FindById(body.ActivityGroupID)

	if err != nil {
		return entityTodo.Todo{}, errors.New("Activity/Not-Found")
	}

	newTodo := entityTodo.Todo{
		ActivityGroupID: body.ActivityGroupID,
		Title:           body.Title,
		IsActive:        body.IsActive,
		Priority:        "very-high",
	}

	return s.repo.Save(newTodo)
}

func (s todoService) UpdateById(id int64, body dtoTodo.UpdateTodoBody) (entityTodo.Todo, error) {
	todo, err := s.repo.FindById(id)

	if err != nil {
		return todo, err
	}

	if body.IsActive != nil {
		todo.IsActive = *body.IsActive
	}

	if body.Priority != nil {
		todo.Priority = *body.Priority
	}

	if body.Title != nil {
		todo.Title = *body.Title
	}

	return s.repo.Save(todo)
}

func (s todoService) DeleteById(id int64) (entityTodo.Todo, error) {
	todo, err := s.repo.FindById(id)

	if err != nil {
		return todo, err
	}

	errDelete := s.repo.Delete(id)

	if errDelete != nil {
		return todo, errDelete
	}

	return todo, nil
}
