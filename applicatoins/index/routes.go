package index

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	index := r.Group("")
	{
		index.GET("", indexPage)
	}
}
