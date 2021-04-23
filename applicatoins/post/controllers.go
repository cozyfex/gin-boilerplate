package post

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func page(c *gin.Context) {
	session := sessions.Default(c)

	c.HTML(http.StatusOK, "post/page", gin.H{
		"session_test": session.Get("test"),
	})
}
