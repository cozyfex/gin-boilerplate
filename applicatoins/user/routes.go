package user

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("", index)
	}
}
