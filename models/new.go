package models

import (
	"errors"
	"fmt"
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

const itemsPerPage = 25

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
func GetNewsItems(db *neoism.Database, tags []string, providers []string, page int) (*[]NewsItem, error) {
	var news []NewsItem

	matchClause := []string{"MATCH (new:NewsItem)"}

	matchClause = append(matchClause, un.MapString(func(tag string) string {
		return fmt.Sprintf("(new:NewsItem)--(:Tag{name: \"%s\"})", strings.TrimSpace(tag))
	}, tags)...)

	match := strings.Join(append(matchClause, "(new:NewsItem)--(p:NewsProvider)"), ", ")

	where := ""
	if len(providers) != 0 {
		names := un.MapString(func(provider string) string {
			return fmt.Sprintf("\"%s\"", strings.TrimSpace(provider))
		}, providers)

		where = fmt.Sprintf("WHERE p.name in [%s]", strings.Join(names, ", "))
	}

	paging := fmt.Sprintf("SKIP %d LIMIT %d", page*itemsPerPage, itemsPerPage)
	query := "RETURN ID(new) as id, new.title as title, new.url as url, new.image as image, new.summary as summary"

	if err := db.Cypher(&neoism.CypherQuery{
		Statement: fmt.Sprintf("%s %s %s %s", match, where, query, paging),
		Result:    &news,
	}); err != nil {
		return nil, err
	}
	return &news, nil
}
