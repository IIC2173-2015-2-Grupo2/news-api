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
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Source   string `json:"source"`
	Location string `json:"location"`
	Language string `json:"language"`
	Body     string `json:"body"`
	Image    string `json:"image"`
	Company  string `json:"company"`
	Person   string `json:"person"`
}

// ---------------------------------------------------------------------------

const itemsPerPage = 25

/*
GetNewsItem returns the new with that id
*/
func GetNewsItem(db *neoism.Database, id int) (*NewsItem, error) {
	var news []NewsItem
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (new:NewsItem)<-[r]-(k:NewsProvider)
								WHERE ID(new) = {id}
								RETURN DISTINCT ID(new) as id, new.title as title, new.url as url,new.image as image, new.body as body, new.language as language, k.name as source`,
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
func GetNewsItems(db *neoism.Database, tags []string, providers []string, categories []string, people []string, locations []string, companies []string, page int) (*[]NewsItem, error) {
	var news []NewsItem

	matchClause := []string{"MATCH (new:NewsItem)<-[r]-(p:NewsProvider)"}

	matchClause = append(matchClause, un.MapString(func(tag string) string {
		return fmt.Sprintf("(new:NewsItem)--(:Tag{name: \"%s\"})", strings.TrimSpace(tag))
	}, tags)...)

	matchClause = append(matchClause, un.MapString(func(category string) string {
		return fmt.Sprintf("(new:NewsItem)--(:Category{name: \"%s\"})", strings.TrimSpace(category))
	}, categories)...)

	// match := strings.Join(append(matchClause, "(new:NewsItem)--(p:NewsProvider)"), ", ")
	// match := strings.Join(matchClause, ", ")

	match := strings.Join(append(matchClause, "(new:NewsItem)"), ", ")

	query := "RETURN DISTINCT ID(new) as id, new.title as title, new.url as url, new.image as image, new.body as body,new.language as language, p.name as source"
	where := ""
	if len(providers) != 0 {
		names := un.MapString(func(provider string) string {
			return fmt.Sprintf("\"%s\"", strings.TrimSpace(provider))
		}, providers)

		where = fmt.Sprintf("WHERE p.name in [%s]", strings.Join(names, ", "))
	}

	if len(locations) != 0 {
		match = match + ", (new:NewsItem)<-[h]-(l:Location)"

		names := un.MapString(func(location string) string {
			return fmt.Sprintf("\"%s\"", strings.TrimSpace(location))
		}, locations)

		where = fmt.Sprintf("WHERE l.name in [%s]", strings.Join(names, ", "))
		query = query + ", l.name as location"

	}

	if len(companies) != 0 {
		match = match + ", (new:NewsItem)--(co:Company)"

		names := un.MapString(func(company string) string {
			return fmt.Sprintf("\"%s\"", strings.TrimSpace(company))
		}, companies)

		where = fmt.Sprintf("WHERE co.name in [%s]", strings.Join(names, ", "))
		query = query + ", co.name as company"
	}

	if len(people) != 0 {
		match = match + ", (new:NewsItem)--(person:Person)"
		names := un.MapString(func(person string) string {
			return fmt.Sprintf("\"%s\"", strings.TrimSpace(person))
		}, people)

		where = fmt.Sprintf("WHERE person.name in [%s]", strings.Join(names, ", "))
		query = query + ", person.name as person"

	}

	paging := fmt.Sprintf("SKIP %d LIMIT %d", page*itemsPerPage, itemsPerPage)

	// fmt.Printf("%s %s %s %s", match, where, query, paging)
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: fmt.Sprintf("%s %s %s %s", match, where, query, paging),
		Result:    &news,
	}); err != nil {
		return nil, err
	}
	return &news, nil
}
