package models

import (
	"errors"
	"fmt"
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
func GetNewsItems(db *neoism.Database, tags []string) (*[]NewsItem, error) {
	var news []NewsItem
	matchClause := ""

	var stat string

	if(tags != nil){
		for _, tag := range tags {
		  // index is the index where we are
		  // tag is the tag from tags for where we are
			matchClause = matchClause + "(new:NewsItem)--(:Tag{name: \"" + tag + "\"}) , "

		}
		matchClause = matchClause[0:len(matchClause) - 2]
	}else{
		matchClause = "(new:NewsItem)"
	}
	stat = "MATCH " + matchClause + 
					 "RETURN new.title as title, new.url as url"
	
	
	fmt.Printf("Match Clause\n")
	fmt.Printf(matchClause+"\n")
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: stat,
		Parameters: neoism.Props{"matchClause": matchClause},
		Result: &news,
	}); err != nil {
		return nil, err
	}
	return &news, nil
}
