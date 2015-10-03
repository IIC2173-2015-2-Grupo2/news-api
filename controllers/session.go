package controllers

import (
	"net/http"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/IIC2173-2015-2-Grupo2/news-api/util"
	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
)

/*
SessionController CRUD
*/
type SessionController struct {
	DB *neoism.Database
}

/*
Token get token
*/
func (n *SessionController) Token(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")

	if user, err := models.FindUserByUsername(n.DB, username); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username not found"})

	} else if err := util.ValidatePass(password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
	}

	claims := map[string]interface{}{
		"ID": username,
	}

	if token, err := util.Token("", claims); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}
