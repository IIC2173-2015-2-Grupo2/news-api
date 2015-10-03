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
	username, password := c.PostForm("username"), c.PostForm("password")

	if duplicated, err := models.FindUserByUsername(n.DB, username); err != nil || duplicated != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username already exists"})

	} else if hashpassword, err := util.HashPass(password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})

	} else {
		name, email := c.PostForm("email"), c.PostForm("name")

		user := models.User{
			Name:     name,
			Username: username,
			Email:    email,
			Password: hashpassword,
		}

		if err := user.Create(n.DB); err != nil {
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
	username, password := c.PostForm("username"), c.PostForm("password")

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
