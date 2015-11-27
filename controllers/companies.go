package controllers

import (
  "net/http"

  "github.com/IIC2173-2015-2-Grupo2/news-api/models"
  "github.com/gin-gonic/gin"
)

/*
CompaniesController CRUD
*/
type CompaniesController struct {
  Neo4jBase
}

/*
Index show list
*/
func (n *CompaniesController) Index(c *gin.Context) {
  if companies, err := models.GetCompanies(n.DB); err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

  } else {
    n.Log("Companies", "Index")
    c.JSON(http.StatusOK, gin.H{"companies": companies})
  }
}
