package controllers

import (
  "net/http"

  "github.com/IIC2173-2015-2-Grupo2/news-api/models"
  "github.com/gin-gonic/gin"
)

/*
CategoriesController CRUD
*/
type CategoriesController struct {
  Neo4jBase
}

/*
Index show list
*/
func (n *CategoriesController) Index(c *gin.Context) {
  if categories, err := models.GetCategories(n.DB); err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

  } else {
    n.Log("Categories", "Index")
    c.JSON(http.StatusOK, gin.H{"categories": categories})
  }
}
