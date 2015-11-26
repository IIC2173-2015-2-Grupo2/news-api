package models

import (
  "github.com/jmcvetta/neoism"
)

/*
Category model
*/
type Category struct {
  ID   int64  `json:"id"`
  Name string `json:"name"`
}

// ---------------------------------------------------------------------------

/*
GetCategorys returns collection of news
*/
func GetCategories(db *neoism.Database) (*[]Category, error) {
  var categories []Category
  if err := db.Cypher(&neoism.CypherQuery{
    Statement: `MATCH (category:Category)
                RETURN ID(category) as id, category.name as name`,
    Result: &categories,
  }); err != nil {
    return nil, err
  }
  return &categories, nil
}
