package activities

import (
	"github.com/gin-gonic/gin"
	handlerActivity "github.com/nach9/go-todolist/pkg/activities/handler"
	repositoryActivity "github.com/nach9/go-todolist/pkg/activities/repository"
	serviceActivity "github.com/nach9/go-todolist/pkg/activities/service"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	activityRepo := repositoryActivity.NewActivityRepo(db)
	activityService := serviceActivity.NewActivityService(activityRepo)
	activityHandler := handlerActivity.NewActivityHandler(activityService)

	routes := r.Group("/activity-groups")

	routes.GET("/", activityHandler.GetActivityList)
	routes.GET("/:id", activityHandler.GetActivityById)
	routes.PATCH("/:id", activityHandler.UpdateActivity)
	routes.POST("/", activityHandler.CreateActivity)
	routes.DELETE("/:id", activityHandler.DeleteActivity)
}
