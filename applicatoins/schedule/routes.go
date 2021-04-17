package schedule

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	schedule := r.Group("/schedule")
	{
		schedule.GET("", index)
	}
}
