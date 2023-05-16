package handlerActivity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dtoActivity "github.com/nach9/go-todolist/pkg/activities/dto"
	serviceActivity "github.com/nach9/go-todolist/pkg/activities/service"
	"github.com/nach9/go-todolist/pkg/common/validation"
)

type ActivityHandler interface {
	GetActivityById(c *gin.Context)
	GetActivityList(c *gin.Context)
	CreateActivity(c *gin.Context)
	UpdateActivity(c *gin.Context)
	DeleteActivity(c *gin.Context)
}

type activityHandler struct {
	service serviceActivity.ActivityService
}

func NewActivityHandler(service serviceActivity.ActivityService) ActivityHandler {
	return &activityHandler{service}
}

func (h activityHandler) GetActivityById(c *gin.Context) {
	id := c.Param("id")

	activityId, _ := strconv.Atoi(id)

	activity, err := h.service.GetById(int64(activityId))

	if err != nil {
		exceptionNotFound(c, activityId)
		return
	}

	c.JSON(http.StatusOK, dtoActivity.ActivityResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})

}

func (h activityHandler) GetActivityList(c *gin.Context) {
	activities := h.service.GetAll()

	c.JSON(http.StatusOK, dtoActivity.ActivityListResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

func (h activityHandler) CreateActivity(c *gin.Context) {
	var body dtoActivity.CreateActivityBody

	if err := c.ShouldBind(&body); err != nil {
		validation.ErrorValidation(err, c, body)
		return
	}

	newActivity, err := h.service.Create(body)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, dtoActivity.ActivityResponse{
		Status:  "Success",
		Message: "Success",
		Data:    newActivity,
	})
}

func (h activityHandler) UpdateActivity(c *gin.Context) {
	var body dtoActivity.UpdateActivityBody

	if err := c.ShouldBind(&body); err != nil {
		validation.ErrorValidation(err, c, body)
		return
	}

	id := c.Param("id")
	activityId, _ := strconv.Atoi(id)

	activity, err := h.service.UpdateById(int64(activityId), body)

	if err != nil {
		exceptionNotFound(c, activityId)
		return
	}

	c.JSON(http.StatusOK, dtoActivity.ActivityResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})

}

func (h activityHandler) DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	activityId, _ := strconv.Atoi(id)

	_, err := h.service.DeleteById(int64(activityId))

	if err != nil {
		exceptionNotFound(c, activityId)
		return
	}

	c.JSON(http.StatusOK, dtoActivity.ActivityDeleteResponse{
		Status:  "Success",
		Message: "Success",
		Data:    &dtoActivity.Blank{},
	})
}
