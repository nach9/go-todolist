package handlerActivity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func exceptionNotFound(c *gin.Context, id int) {
	c.AbortWithStatusJSON(http.StatusNotFound,
		gin.H{
			"status":  "Not Found",
			"message": "Activity with ID " + strconv.Itoa(id) + " Not Found",
		})
}
