package handlerActivity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nach9/go-todolist/pkg/common/validation"
	dtoTodo "github.com/nach9/go-todolist/pkg/todos/dto"
	serviceTodo "github.com/nach9/go-todolist/pkg/todos/service"
)

type TodoHandler interface {
	GetTodoById(c *gin.Context)
	GetTodoList(c *gin.Context)
	CreateTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type todoHandler struct {
	service serviceTodo.TodoService
}

func NewTodoHandler(service serviceTodo.TodoService) TodoHandler {
	return &todoHandler{service}
}

func (h todoHandler) GetTodoById(c *gin.Context) {
	id := c.Param("id")

	todoId, _ := strconv.Atoi(id)

	todo, err := h.service.GetById(int64(todoId))

	if err != nil {
		exceptionNotFound(c, todoId)
		return
	}

	c.JSON(http.StatusOK, dtoTodo.TodoResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})

}

func (h todoHandler) GetTodoList(c *gin.Context) {
	var param dtoTodo.GetAllParam

	if err := c.ShouldBind(&param); err != nil {
		validation.ErrorValidation(err, c, param)
		return
	}

	todos := h.service.GetAll(param)

	c.JSON(http.StatusOK, dtoTodo.TodoListResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

func (h todoHandler) CreateTodo(c *gin.Context) {
	var body dtoTodo.CreateTodoBody

	if err := c.ShouldBind(&body); err != nil {
		validation.ErrorValidation(err, c, body)
		return
	}

	newTodo, err := h.service.Create(body)

	if err != nil {
		if err.Error() == "Activity/Not-Found" {
			exceptionActivityNotFound(c, int(body.ActivityGroupID))
			return
		}

		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, dtoTodo.TodoResponse{
		Status:  "Success",
		Message: "Success",
		Data:    newTodo,
	})
}

func (h todoHandler) UpdateTodo(c *gin.Context) {
	var body dtoTodo.UpdateTodoBody

	if err := c.ShouldBind(&body); err != nil {
		validation.ErrorValidation(err, c, body)
		return
	}

	id := c.Param("id")
	todoId, _ := strconv.Atoi(id)

	todo, err := h.service.UpdateById(int64(todoId), body)

	if err != nil {
		exceptionNotFound(c, todoId)
		return
	}

	c.JSON(http.StatusOK, dtoTodo.TodoResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})

}

func (h todoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	todoId, _ := strconv.Atoi(id)

	_, err := h.service.DeleteById(int64(todoId))

	if err != nil {
		exceptionNotFound(c, todoId)
		return
	}

	c.JSON(http.StatusOK, dtoTodo.TodoDeleteResponse{
		Status:  "Success",
		Message: "Success",
		Data:    &dtoTodo.Blank{},
	})
}
