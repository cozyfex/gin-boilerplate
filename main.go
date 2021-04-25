package main

import (
	"gin-boilerplate/applicatoins/index"
	"gin-boilerplate/applicatoins/post"
	"gin-boilerplate/applicatoins/schedule"
	"gin-boilerplate/applicatoins/user"
	"gin-boilerplate/libs/loadtemplates"
	cfMySQLStore "gin-boilerplate/libs/mysqlstore"
	"gin-boilerplate/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	errGoDotEnv := godotenv.Load()
	if errGoDotEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// Get env values
	dsn := os.Getenv("dsn")
	sessionTable := os.Getenv("sessionTable")
	sessionPath := os.Getenv("sessionPath")
	sessionMaxAge, _ := strconv.Atoi(os.Getenv("sessionMaxAge"))
	sessionSecretKey := os.Getenv("sessionSecretKey")

	// MySQL session
	store, errStore := cfMySQLStore.NewStore(dsn, sessionTable, sessionPath, sessionMaxAge, []byte(sessionSecretKey))
	if errStore != nil {
		panic(errStore)
	}

	// Gorm
	db, errGorm := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errGorm != nil {
		panic(errGorm)
	}
	models.Migrate(db)

	// Create route
	r := gin.Default()
	// Session
	r.Use(sessions.Sessions("CF_SESSION", store))
	// Static files
	r.Static("/static", "./assets")
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
