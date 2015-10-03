package controllers

import (
	"net/http"
	"strconv"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
)

/*
UsersController CRUD
*/
type UsersController struct {
	DB *neoism.Database
}

/*
Index show list
*/
func (n *UsersController) Index(c *gin.Context) {
	if users, err := models.GetUsers(n.DB); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

/*
Show specific user
*/
func (n *UsersController) Show(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else if user, err := models.GetUser(n.DB, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
