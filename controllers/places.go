package controllers

import (
	"net/http"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/gin-gonic/gin"
)

/*
PlacesController CRUD
*/
type PlacesController struct {
	Base
}

/*
Index show list
*/
func (n *PlacesController) Index(c *gin.Context) {
	if places, err := models.GetPlaces(n.DB); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	} else {
		n.Log("Places", "Index")
		c.JSON(http.StatusOK, gin.H{"places": places})
	}
}
