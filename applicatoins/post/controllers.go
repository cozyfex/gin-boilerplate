package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func page(c *gin.Context) {
	c.HTML(http.StatusOK, "post/page", gin.H{
		"title": "Home Page",
	})
}
