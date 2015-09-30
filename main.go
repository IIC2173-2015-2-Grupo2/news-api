package main

import (
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
	var err error
	if environment := os.Getenv("ENVIRONMENT"); environment == "PRODUCTION" {
		db = database.Connect(
			os.Getenv("NEO4USER"),
			os.Getenv("NEO4PASSWORD"),
			os.Getenv("NEO4JHOST"),
			os.Getenv("NEO4JPORT"),
		)
	}else{
		// db, _ = neoism.Connect("http://localhost:7474/db/data")
		db, err = neoism.Connect("http://neo4j:12345678@localhost:7474/db/data")
		print(err)
// ://neo4j:password@server:7474/db/data
		

	}

	// Router
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

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
	// Secret hash
	secret := os.Getenv("SECRET_HASH")

	// Controllers
	newsController := controllers.NewsController{DB: db}

	// Public API
	public := router.Group("/api/v1")
	public.GET("/auth", func(c *gin.Context) {
		claims := map[string]interface{}{
			"ID": "Patiwi",
		}
		if token, err := middleware.Token("", claims); err != nil {
			c.JSON(500, gin.H{"error": err})
		} else {
			c.JSON(200, gin.H{"token": token})
		}
	})
	public.GET("/news", newsController.PublicIndex)
	// Private API
	private := router.Group("/api/v1/private")
	private.Use(middleware.AuthMiddleware(secret))

	private.GET("/news", newsController.Index)
	private.GET("/news/:id", newsController.Show)
}
