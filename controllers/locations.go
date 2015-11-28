package controllers

import (
	"net/http"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/gin-gonic/gin"
)

/*
LocationsController CRUD
*/
type LocationsController struct {
	Neo4jBase
}

/*
Index show list
*/
func (n *LocationsController) Index(c *gin.Context) {
	if locations, err := models.GetLocations(n.DB); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	} else {
		n.Log("Locations", "Index")
		c.JSON(http.StatusOK, gin.H{"locations": locations})
	}
}
