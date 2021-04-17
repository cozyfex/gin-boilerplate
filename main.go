package main

import (
	"gin-boilerplate/applicatoins/index"
	"gin-boilerplate/applicatoins/post"
	"gin-boilerplate/applicatoins/schedule"
	"gin-boilerplate/applicatoins/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/**/*")

	index.Routes(r)
	post.Routes(r)
	user.Routes(r)
	schedule.Routes(r)

	err := r.Run()
	if err != nil {
		return
	}
}
