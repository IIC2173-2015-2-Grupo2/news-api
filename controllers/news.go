package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
)

/*
NewsController CRUD
*/
type NewsController struct {
	DB *neoism.Database
}

/*
Index show list
*/
func (n *NewsController) Index(c *gin.Context) {
	c.JSON(200, gin.H{"Controller": "News", "Action": "Index"})
}

/*
Show specific new
*/
func (n *NewsController) Show(c *gin.Context) {
	// Docs: https://github.com/gin-gonic/gin#xml-and-json-rendering
	var msg struct {
		ID   string
		Tag  string
		Date string
	}
	msg.ID = c.Params.ByName("id")
	msg.Tag = c.Query("tag") // default: ""
	msg.Date = c.DefaultQuery("date", "2015-05-25")

	c.JSON(200, gin.H{"Controller": "News", "Action": "Show", "Message": msg})

	/* http://localhost:8000/api/v1/news/1
	 * Output:
			{
			  "Action": "Show",
			  "Controller": "News",
			  "Message": {
			    "ID": "1",
			    "Tag": "",
			    "Date": "2015-05-25"
			  }
			}
	*/
}
