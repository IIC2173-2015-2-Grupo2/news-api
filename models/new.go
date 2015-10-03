package models

import (
	"errors"

	"github.com/jmcvetta/neoism"
)

/*
New model
*/
type New struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// ---------------------------------------------------------------------------

/*
GetNew returns the new with that id
*/
func GetNew(db *neoism.Database, id int) (*New, error) {
	var news []New
	if err := db.Cypher(&neoism.CypherQuery{
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
func GetNews(db *neoism.Database) (*[]New, error) {
	var news []New
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (new:New)
								RETURN new.title as title, new.url as url`,
		Result: &news,
	}); err != nil {
		return nil, err
	}
	return &news, nil
}
