package handlerActivity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type exceptionMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func exceptionNotFound(c *gin.Context, id int) {
	c.AbortWithStatusJSON(http.StatusNotFound, exceptionMessage{
		Status:  "Not Found",
		Message: "Todo with ID " + strconv.Itoa(id) + " Not Found",
	})
}

func exceptionActivityNotFound(c *gin.Context, id int) {
	c.AbortWithStatusJSON(http.StatusNotFound, exceptionMessage{
		Status:  "Not Found",
		Message: "Activity with ID " + strconv.Itoa(id) + " Not Found",
	})
}
