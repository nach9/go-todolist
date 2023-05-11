package repositoryTodo

import (
	dtoTodo "github.com/nach9/go-todolist/pkg/todos/dto"
	entityTodo "github.com/nach9/go-todolist/pkg/todos/entity"
	"gorm.io/gorm"
)

type TodoRepo interface {
	FindAll(param dtoTodo.GetAllParam) []entityTodo.Todo
	FindById(id int64) (entityTodo.Todo, error)
	Save(todo entityTodo.Todo) (entityTodo.Todo, error)
	Delete(id int64) error
}

type todoRepo struct {
	DB *gorm.DB
}

func NewTodoRepo(DB *gorm.DB) TodoRepo {
	return &todoRepo{DB}
}

func (r *todoRepo) FindAll(param dtoTodo.GetAllParam) []entityTodo.Todo {
	var todos []entityTodo.Todo

	query := r.DB

	if param.ActivityGroupID != nil {
		query = query.
			Where("activity_group_id=?", *param.ActivityGroupID)
	}

	query.Find(&todos)

	return todos
}

func (r *todoRepo) FindById(id int64) (entityTodo.Todo, error) {
	var todo entityTodo.Todo
	result := r.DB.First(&todo, id)

	if result.Error != nil {
		return todo, result.Error
	}

	return todo, nil
}

func (r *todoRepo) Save(todo entityTodo.Todo) (entityTodo.Todo, error) {
	result := r.DB.Save(&todo)

	if result.Error != nil {
		return todo, result.Error
	}

	return todo, nil
}

func (r *todoRepo) Delete(id int64) error {
	var todo entityTodo.Todo

	result := r.DB.Delete(&todo, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
