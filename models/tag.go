package models

import (
	"github.com/jmcvetta/neoism"
)

/*
Tag model
*/
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ---------------------------------------------------------------------------

/*
GetTags returns collection of news
*/
func GetTags(db *neoism.Database) (*[]Tag, error) {
	var tags []Tag
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (tag:Tag)
                RETURN ID(tag) as id, tag.name as name`,
		Result: &tags,
	}); err != nil {
		return nil, err
	}
	return &tags, nil
}
