package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index", gin.H{
		"title": "Main",
	})
}
