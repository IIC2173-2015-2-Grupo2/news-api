package models

import (
	"errors"
	"strings"

	"github.com/jmcvetta/neoism"
	un "github.com/tobyhede/go-underscore"
)

/*
NewsItem model
*/
type NewsItem struct {
	ID      int64  `json:"id"`
	TITLE   string `json:"title"`
	URL     string `json:"url"`
	SUMMARY string `json:"summary"`
	IMAGE   string `json:"image"`
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
								RETURN ID(new) as id, new.title as title, new.url as url,new.image as image, new.summary as summary`,
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
func GetNewsItems(db *neoism.Database, tags []string, providers []string) (*[]NewsItem, error) {
	var news []NewsItem

	matchClause := []string{"MATCH (new:NewsItem)"}

	matchClause = append(matchClause, un.MapString(func(tag string) string {
		return "(new:NewsItem)--(:Tag{name: \"" + strings.TrimSpace(tag) + "\"})"
	}, tags)...)

	match := strings.Join(append(matchClause, "(new:NewsItem)--(p:NewsProvider)"), ", ")

	where := ""
	if len(providers) != 0 {
		where = "WHERE p.name in [" + strings.Join(un.MapString(func(provider string) string {
			return "\"" + strings.TrimSpace(provider) + "\""
		}, providers), ", ") + "]"
	}

	if err := db.Cypher(&neoism.CypherQuery{
		Statement: match + " " + where + "RETURN ID(new) as id, new.title as title, new.url as url, new.image as image, new.summary as summary",
		Result:    &news,
	}); err != nil {
		return nil, err
	}
	return &news, nil
}
