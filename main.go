package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"

	"github.com/IIC2173-2015-2-Grupo2/news-api/controllers"
	"github.com/IIC2173-2015-2-Grupo2/news-api/database"
	"github.com/IIC2173-2015-2-Grupo2/news-api/middleware"
)

const (
	port = ":8000"
)

func main() {
	// Database setup
	var db *neoism.Database
	if os.Getenv("NEO4J_HOST") != "" && os.Getenv("NEO4J_PORT") != "" {
		if connected, err := database.Connect(
			os.Getenv("NEO4J_USER"),
			os.Getenv("NEO4J_PASS"),
			os.Getenv("NEO4J_HOST"),
			os.Getenv("NEO4J_PORT"),
		); err != nil {
			log.Fatal(err)
		} else {
			db = connected
		}
	}

	// Setup environment
	if environment := os.Getenv("ENVIRONMENT"); environment == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		if db != nil {
			db.Session.Log = true
		}
	}

	// Start
	Server(db).Run(port)
}

/*
Server server
*/
func Server(db *neoism.Database) *gin.Engine {

	// Router
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.GZIP())

	// Welcome
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/v1")
	})

	// Configure API v1
	apiv1(router, db)

	return router
}

func apiv1(router *gin.Engine, db *neoism.Database) {
	secret := os.Getenv("SECRET_HASH")

	// Controllers --------------------------------------------------------------
	newsController := controllers.NewsController{DB: db}
	usersController := controllers.UsersController{DB: db}
	sessionController := controllers.SessionController{DB: db, SecretHash: secret}
	// --------------------------------------------------------------------------

	// Public API ---------------------------------------------------------------
	public := router.Group("/api/v1")

	public.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "on"})
	})
	// --------------------------------------------------------------------------

	// Auth API -----------------------------------------------------------------
	auth := router.Group("/api/v1/auth")

	auth.POST("/token", sessionController.Authorize)
	auth.POST("/signup", sessionController.Create)
	// --------------------------------------------------------------------------

	// Private API --------------------------------------------------------------
	private := router.Group("/api/v1/private")

	if auth := os.Getenv("AUTH"); auth == "ENABLE" {
		private.Use(middleware.APIToken())
		private.Use(middleware.JWTAuth(secret))
	}

	private.GET("/news", newsController.Index)
	private.GET("/search", newsController.Search)
	private.GET("/news/:id", newsController.Show)
	private.GET("/users", usersController.Index)
	private.GET("/users/:id", usersController.Show)
	// --------------------------------------------------------------------------
}
