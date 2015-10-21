package models

import (
	"errors"
	"fmt"
	"strings"
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
func GetNewsItems(db *neoism.Database, tags []string, providers []string) (*[]NewsItem, error) {

	var news []NewsItem
	matchClause := "MATCH (new:NewsItem)"
	whereClause := ""


	var stat string

	if(tags != nil){
		matchClause= matchClause+", "
		for _, tag := range tags {
		  // index is the index where we are
		  // tag is the tag from tags for where we are
			matchClause = matchClause + "(new:NewsItem)--(:Tag{name: \"" + strings.TrimSpace(tag) + "\"}) , "

		}
		matchClause = matchClause[0:len(matchClause) - 2]
	}
	if(providers != nil){
		whereClause= "WHERE "
		matchClause = matchClause + ", (new:NewsItem)--(p:NewsProvider)"
		arrayStr := "["
		for _, provider := range providers {
		  // index is the index where we are
		  // provider is the provider from providers for where we are
			arrayStr = arrayStr + "\""+strings.TrimSpace(provider)+"\", "

		}
		arrayStr=arrayStr+"\"\"]"
		whereClause = "WHERE p.name in "+arrayStr
	}

	stat =  matchClause + " "+
					whereClause + " "+
					 "RETURN new.title as title, new.url as url"
	
	
	fmt.Printf("Match Clause\n")
	fmt.Printf(matchClause+"\n")
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: stat,
		Result: &news,
	}); err != nil {
		return nil, err
	}
	return &news, nil
}
