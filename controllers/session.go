package controllers

import (
	"net/http"

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
	user, _ := c.PostForm("username"), c.PostForm("password")

	claims := map[string]interface{}{
		"ID": user,
	}

	if token, err := util.Token("", claims); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}
