package schedule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "schedule/index", gin.H{
		"title": "Schedule Index",
	})
}
