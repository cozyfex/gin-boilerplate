package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index", gin.H{
		"title": "User Index",
	})
}
