package controllers

import (
	"net/http"
	"strconv"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
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
	if news, err := models.GetNewsItems(n.DB); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"news": news})
	}
}

/*
Search a new
*/
func (n *NewsController) Search(c *gin.Context) {
	// TODO: search
	n.Index(c)
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
		c.JSON(http.StatusOK, gin.H{"new": new})
	}
}
