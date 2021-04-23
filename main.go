package main

import (
	"gin-boilerplate/applicatoins/index"
	"gin-boilerplate/applicatoins/post"
	"gin-boilerplate/applicatoins/schedule"
	"gin-boilerplate/applicatoins/user"
	"gin-boilerplate/libs/loadtemplates"
	cfMySQLStore "gin-boilerplate/libs/mysqlstore"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get env values
	dsn := os.Getenv("dsn")
	sessionTable := os.Getenv("sessionTable")
	sessionPath := os.Getenv("sessionPath")
	sessionMaxAge, _ := strconv.Atoi(os.Getenv("sessionMaxAge"))
	sessionSecretKey := os.Getenv("sessionSecretKey")

	// MySQL session
	store, err := cfMySQLStore.NewStore(dsn, sessionTable, sessionPath, sessionMaxAge, []byte(sessionSecretKey))
	if err != nil {
		panic(err)
	}

	// Create route
	r := gin.Default()
	// Session
	r.Use(sessions.Sessions("CF_SESSION", store))
	// Static files
	r.Static("/assets", "./assets")
	// Load templates
	files := loadtemplates.GetFiles("templates")
	r.LoadHTMLFiles(files...)

	// Register routes
	index.Routes(r)
	post.Routes(r)
	user.Routes(r)
	schedule.Routes(r)

	// Run route
	routeError := r.Run()
	if routeError != nil {
		panic(routeError)
	}
}
