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
	DB         *neoism.Database
	SecretHash string
}

/*
Create account
*/
func (n *SessionController) Create(c *gin.Context) {
	var user models.User
	c.Bind(&user)

	if len(user.Username) == 0 || len(user.Password) == 0 || len(user.Name) == 0 || len(user.Email) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing fields"})

	} else if duplicated, err := models.FindUserByUsername(n.DB, user.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})

	} else if duplicated != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username already exists"})

	} else if hashpassword, err := util.HashPass(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})

	} else {
		user.Password = hashpassword

		if _, err := user.Save(n.DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})

		} else {
			n.Token(c, &user)
		}
	}
}

/*
Authorize user to access to private resources
*/
func (n *SessionController) Authorize(c *gin.Context) {
	username, password := c.Param("username"), c.Param("password")

	if user, err := models.FindUserByUsername(n.DB, username); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username not found"})

	} else if err := util.ValidatePass(password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})

	} else {
		n.Token(c, user)
	}
}

/*
Token creation for an user
*/
func (n *SessionController) Token(c *gin.Context, user *models.User) {
	claims := map[string]interface{}{
		"ID": user.Username,
	}

	if token, err := util.Token(n.SecretHash, claims); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})

	} else {
		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}
