package controllers

import (
  "net/http"

  "github.com/IIC2173-2015-2-Grupo2/news-api/models"
  "github.com/gin-gonic/gin"
)

/*
PeopleController CRUD
*/
type PeopleController struct {
  Neo4jBase
}

/*
Index show list
*/
func (n *PeopleController) Index(c *gin.Context) {
  if people, err := models.GetPeople(n.DB); err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

  } else {
    n.Log("People", "Index")
    c.JSON(http.StatusOK, gin.H{"people": people})
  }
}
