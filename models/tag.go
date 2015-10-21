package models

import (

  "github.com/jmcvetta/neoism"
)

/*
Tag model
*/
type Tag struct {
  Name string `json:"tag"`
}

// ---------------------------------------------------------------------------

/*
GetNewsItems returns collection of news
*/
func GetTags(db *neoism.Database) (*[]Tag, error) {
  var tags []Tag
  if err := db.Cypher(&neoism.CypherQuery{
    Statement: `MATCH (tag:Tag)
                RETURN tag.name as tag`,
    Result: &tags,
  }); err != nil {
    return nil, err
  }
  return &tags, nil
}
