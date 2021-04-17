package post

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	post := r.Group("/post")
	{
		post.GET("/page", page)
	}
}
