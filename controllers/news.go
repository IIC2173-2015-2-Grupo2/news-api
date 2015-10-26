package controllers

import (
	"net/http"
	"strconv"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/gin-gonic/gin"
)

/*
NewsController CRUD
*/
type NewsController struct {
	Base
}

/*
Index show list
*/
func (n *NewsController) Index(c *gin.Context) {
	if news, err := models.GetNewsItems(n.DB, nil, nil); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
	} else {
		n.Log("NewsItems", "Index")
		c.JSON(http.StatusOK, gin.H{"news": news})
	}
}

/*
Search a new
*/
func (n *NewsController) Search(c *gin.Context) {
	tags := c.Request.URL.Query()["tags"]
	providers := c.Request.URL.Query()["providers"]

	if news, err := models.GetNewsItems(n.DB, tags, providers); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	} else {
		n.Log("NewsItems", "Search")
		c.JSON(http.StatusOK, gin.H{"news": news})
	}
}

/*
Show specific new
*/
func (n *NewsController) Show(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else if new, err := models.GetNewsItem(n.DB, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	} else {
		n.Log("NewsItems", c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"new": new})
	}
}
