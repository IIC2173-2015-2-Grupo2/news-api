package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jmcvetta/neoism"
	"github.com/jpillora/go-ogle-analytics"
	_ "github.com/lib/pq"

	"github.com/IIC2173-2015-2-Grupo2/news-api/controllers"
	"github.com/IIC2173-2015-2-Grupo2/news-api/database"
	"github.com/IIC2173-2015-2-Grupo2/news-api/middleware"
	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
)

const (
	port = ":8000"
)

func main() {
	// Neo4J Database setup
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

	// Postgres Database setup
	var pgdb *gorm.DB
	if connected, err := gorm.Open("postgres", "user=postgres dbname=newsapi sslmode=disable host=db"); err != nil {
		// if connected, err := gorm.Open("postgres", "user=newsapi dbname=newsapi sslmode=disable"); err != nil {
		log.Fatal(err)
	} else {
		pgdb = &connected
		pgdb.AutoMigrate(&models.User{})
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
	Server(db, pgdb, analytics).Run(port)
}

/*
Server server
*/
func Server(db *neoism.Database, pgdb *gorm.DB, analytics *ga.Client) *gin.Engine {

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
	apiv1(router, db, pgdb, analytics)

	return router
}

func apiv1(router *gin.Engine, db *neoism.Database, pgdb *gorm.DB, analytics *ga.Client) {
	secret := os.Getenv("SECRET_HASH")

	// Controllers --------------------------------------------------------------
	base := controllers.Base{Analytics: analytics}
	neo4jBaseController := controllers.Neo4jBase{DB: db, Base: base}
	pgBaseController := controllers.PgBase{DB: pgdb, Base: base}

	newsController := controllers.NewsController{Neo4jBase: neo4jBaseController}
	peopleController := controllers.PeopleController{Neo4jBase: neo4jBaseController}
	categoriesController := controllers.CategoriesController{Neo4jBase: neo4jBaseController}
	locationsController := controllers.LocationsController{Neo4jBase: neo4jBaseController}
	tagsController := controllers.TagsController{Neo4jBase: neo4jBaseController}
	companiesController := controllers.CompaniesController{Neo4jBase: neo4jBaseController}
	newsProvidersController := controllers.NewsProvidersController{Neo4jBase: neo4jBaseController}
	usersController := controllers.UsersController{PgBase: pgBaseController}
	sessionController := controllers.SessionController{PgBase: pgBaseController, SecretHash: secret}
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
	private.GET("/people", peopleController.Index)
	private.GET("/locations", locationsController.Index)
	private.GET("/categories", categoriesController.Index)
	private.GET("/companies", companiesController.Index)
	private.GET("/news_providers", newsProvidersController.Index)
	private.GET("/news", newsController.Index)
	private.GET("/search", newsController.Search)
	private.GET("/news/:id", newsController.Show)
	private.GET("/users", usersController.Index)
	private.GET("/users/:id", usersController.Show)

	// --------------------------------------------------------------------------

}
