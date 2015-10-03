package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
)

/*
NewsController CRUD
*/
type NewsController struct {
	DB *neoism.Database
}

/*
GetNew returns the new with that id
*/
func (n *NewsController) GetNew(id int) (*models.New, error) {
	var news []models.New
	if err := n.DB.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (new:New)
								WHERE ID(new) = {id}
								RETURN new.title as title, new.url as url`,
		Parameters: neoism.Props{"id": id},
		Result:     &news,
	}); err != nil {
		return nil, err

	} else if len(news) == 0 {
		return nil, errors.New("not found")

	} else {
		return &news[0], nil
	}
}

/*
GetNews returns collection of news
*/
func (n *NewsController) GetNews() (*[]models.New, error) {
	var news []models.New
	if err := n.DB.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (new:New)
								RETURN new.title as title, new.url as url`,
		Result: &news,
	}); err != nil {
		return nil, err
	}
	return &news, nil
}

/*
Index show list
*/
func (n *NewsController) Index(c *gin.Context) {
	if news, err := n.GetNews(); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"news": news})
	}
}

/*
Show specific new
*/
func (n *NewsController) Show(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else if new, err := n.GetNew(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"new": new})
	}
}
