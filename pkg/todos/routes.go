package todos

import (
	"github.com/gin-gonic/gin"
	repositoryActivity "github.com/nach9/go-todolist/pkg/activities/repository"
	handlerTodo "github.com/nach9/go-todolist/pkg/todos/handler"
	repositoryTodo "github.com/nach9/go-todolist/pkg/todos/repository"
	serviceTodo "github.com/nach9/go-todolist/pkg/todos/service"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	todoRepo := repositoryTodo.NewTodoRepo(db)
	activityRepo := repositoryActivity.NewActivityRepo(db)

	todoService := serviceTodo.NewTodoService(todoRepo, activityRepo)
	todoHandler := handlerTodo.NewTodoHandler(todoService)

	routes := r.Group("/todo-items")

	routes.GET("/", todoHandler.GetTodoList)
	routes.GET("/:id", todoHandler.GetTodoById)
	routes.PATCH("/:id", todoHandler.UpdateTodo)
	routes.POST("/", todoHandler.CreateTodo)
	routes.DELETE("/:id", todoHandler.DeleteTodo)
}
