package controllers

import (
  "net/http"

  "github.com/IIC2173-2015-2-Grupo2/news-api/models"
  "github.com/gin-gonic/gin"
  "github.com/jmcvetta/neoism"
)

/*
TagsController CRUD
*/
type TagsController struct {
  DB *neoism.Database
}

/*
Index show list
*/
func (n *TagsController) Index(c *gin.Context) {


  if tags, err := models.GetTags(n.DB); err != nil {
    c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})

  } else {
    c.JSON(http.StatusOK, gin.H{"tags": tags})
  }
}
