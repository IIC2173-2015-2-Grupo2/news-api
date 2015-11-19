package controllers

import (
	"net/http"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/gin-gonic/gin"
)

/*
NewsProvidersController CRUD
*/
type NewsProvidersController struct {
	Neo4jBase
}

/*
Index show list
*/
func (n *NewsProvidersController) Index(c *gin.Context) {
	if tags, err := models.GetNewsProviders(n.DB); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	} else {
		n.Log("NewsProviders", "Index")
		c.JSON(http.StatusOK, gin.H{"news_providers": tags})
	}
}
