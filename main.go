package main

import (
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func main() {
	app := gin.Default()
	app.GET("/", index)
	app.Run(":8000")
}
