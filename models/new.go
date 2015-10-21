package models

import (
	"errors"

	"github.com/jmcvetta/neoism"
)

/*
NewsItem model
*/
type NewsItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// ---------------------------------------------------------------------------

/*
GetNewsItem returns the new with that id
*/
func GetNewsItem(db *neoism.Database, id int) (*NewsItem, error) {
	var news []NewsItem
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (new:NewsItem)
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
GetNewsItems returns collection of news
*/
func GetNewsItems(db *neoism.Database, tags string[]) (*[]NewsItem, error) {
	var news []NewsItem
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (new:NewsItem)
								RETURN new.title as title, new.url as url`,
		Result: &news,
	}); err != nil {
		return nil, err
	}
	return &news, nil
}
