package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"

	"github.com/IIC2173-2015-2-Grupo2/news-api/controllers"
	"github.com/IIC2173-2015-2-Grupo2/news-api/middleware"
)

var db *neoism.Database
var router *gin.Engine

func main() {
	if environment := os.Getenv("ENVIRONMENT"); environment == "PRODUCTION" {
		db = Connect(
			os.Getenv("NEO4USER"),
			os.Getenv("NEO4PASSWORD"),
			os.Getenv("NEO4JHOST"),
			os.Getenv("NEO4JPORT"),
		)
	}

	router = Router(
		middleware.CORSMiddleware(),
	)

	router.Run(":8000")
}

/*
Connect to database
*/
func Connect(user, password, host, port string) *neoism.Database {
	uri := "http://" + user + ":" + password + "@" + host + ":" + port + "/db/data"
	db, err := neoism.Connect(uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

/*
Router for routing HTPP
*/
func Router(middleware ...gin.HandlerFunc) *gin.Engine {
	// Controllers
	newsController := controllers.NewsController{db}

	// Routing
	router := gin.Default()
	router.Use(middleware...)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/news", newsController.Index)
		v1.GET("/news/:id", newsController.Show)
	}
	router.GET("/", index)

	return router
}

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}
