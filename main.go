package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"

	"github.com/IIC2173-2015-2-Grupo2/news-api/controllers"
	"github.com/IIC2173-2015-2-Grupo2/news-api/database"
	"github.com/IIC2173-2015-2-Grupo2/news-api/middleware"
)

/*
Start point
*/
func main() {
	// Database setup
	var db *neoism.Database
	if environment := os.Getenv("ENVIRONMENT"); environment == "PRODUCTION" {
		if connected, err := database.Connect(
			os.Getenv("NEO4USER"),
			os.Getenv("NEO4PASSWORD"),
			os.Getenv("NEO4JHOST"),
			os.Getenv("NEO4JPORT"),
		); err != nil {
			log.Fatal(err)
		} else {
			db = connected
		}
	}

	// Router
	router := gin.Default()
	router.Use(middleware.CORS())

	// Welcome
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"news-api": "working"})
	})

	// Configure API v1
	apiv1(router, db)

	// Start
	router.Run(":8000")
}

func apiv1(router *gin.Engine, db *neoism.Database) {
	secret := os.Getenv("SECRET_HASH")

	// Controllers --------------------------------------------------------------
	newsController := controllers.NewsController{DB: db}
	sessionController := controllers.SessionController{DB: db}
	// --------------------------------------------------------------------------

	// Public API ---------------------------------------------------------------
	public := router.Group("/api/v1")
	public.GET("/hello")
	// --------------------------------------------------------------------------

	// Auth API
	auth := router.Group("/api/v1/auth")
	auth.POST("/token", sessionController.Token)
	// --------------------------------------------------------------------------

	// Private API --------------------------------------------------------------
	private := router.Group("/api/v1/private")
	private.Use(middleware.JWTAuth(secret))

	private.GET("/news", newsController.Index)
	private.GET("/news/:id", newsController.Show)
	// --------------------------------------------------------------------------
}
