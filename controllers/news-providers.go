package controllers

import (
	"net/http"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
)

/*
NewsProvidersController CRUD
*/
type NewsProvidersController struct {
	DB *neoism.Database
}

/*
Index show list
*/
func (n *NewsProvidersController) Index(c *gin.Context) {
	if tags, err := models.GetNewsProviders(n.DB); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"news_providers": tags})
	}
}
