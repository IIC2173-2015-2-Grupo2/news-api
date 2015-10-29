package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
	"github.com/jpillora/go-ogle-analytics"

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

	// Setup analytics client
	var analytics *ga.Client
	if token := os.Getenv("ANALYTICS_TOKEN"); token == "" {
		fmt.Printf("Analytics token not provided.\n")
	} else if client, err := ga.NewClient(token); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Analytics activated.\n")
		analytics = client
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
	Server(db, analytics).Run(port)
}

/*
Server server
*/
func Server(db *neoism.Database, analytics *ga.Client) *gin.Engine {

	// Router
	router := gin.Default()

	// Test with loader.io
	if token := os.Getenv("LOADER_IO_TOKEN"); token == "" {
		fmt.Printf("Loader io token not provided.\n")
	} else {
		fmt.Printf("Loader io activated.\n")
		router.GET("/"+token, func(c *gin.Context) {
			c.Writer.Header().Set("Content-disposition", "attachment;filename="+token+".txt")
			c.Data(http.StatusOK, "text/plain", []byte(token))
		})
	}

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.GZIP())

	// Welcome
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/v1")
	})

	// Configure API v1
	apiv1(router, db, analytics)

	return router
}

func apiv1(router *gin.Engine, db *neoism.Database, analytics *ga.Client) {
	secret := os.Getenv("SECRET_HASH")

	// Controllers --------------------------------------------------------------
	baseController := controllers.Base{DB: db, Analytics: analytics}
	newsController := controllers.NewsController{Base: baseController}
	tagsController := controllers.TagsController{Base: baseController}
	newsProvidersController := controllers.NewsProvidersController{Base: baseController}
	usersController := controllers.UsersController{Base: baseController}
	sessionController := controllers.SessionController{Base: baseController, SecretHash: secret}
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

	private.GET("/tags", tagsController.Index)
	private.GET("/news_providers", newsProvidersController.Index)
	private.GET("/news", newsController.Index)
	private.GET("/search", newsController.Search)
	private.GET("/news/:id", newsController.Show)
	private.GET("/users", usersController.Index)
	private.GET("/users/:id", usersController.Show)
	// --------------------------------------------------------------------------

}
