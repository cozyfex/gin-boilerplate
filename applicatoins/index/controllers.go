package index

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexPage(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("test", "test134")
	_ = session.Save()

	c.HTML(http.StatusOK, "index", gin.H{
		"title":        "Main",
		"session_test": session.Get("test"),
	})
}
