package controllers

import (
	"net/http"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/gin-gonic/gin"
)

/*
TagsController CRUD
*/
type TagsController struct {
	Base
}

/*
Index show list
*/
func (n *TagsController) Index(c *gin.Context) {
	if tags, err := models.GetTags(n.DB); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	} else {
		n.Log("Tags", "Index")
		c.JSON(http.StatusOK, gin.H{"tags": tags})
	}
}
