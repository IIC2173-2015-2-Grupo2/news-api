package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
)

func connect(user, password, host, port string) *neoism.Database {
	uri := "http://" + user + ":" + password + "@" + host + ":" + port + "/db/data"
	db, err := neoism.Connect(uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func main() {
	db := connect(os.Getenv("NEO4USER"), os.Getenv("NEO4PASSWORD"), os.Getenv("NEO4JHOST"), os.Getenv("NEO4JPORT"))
	fmt.Println("Connected to:", db.Url)

	app := gin.Default()
	app.GET("/", index)
	app.Run(":8000")
}
