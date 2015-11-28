package models

import (
  "github.com/jmcvetta/neoism"
)

/*
Location model
*/
type Location struct {
  ID   int64  `json:"id"`
  Name string `json:"name"`
}

// ---------------------------------------------------------------------------

/*
GetLocations returns collection of news
*/
func GetLocations(db *neoism.Database) (*[]Location, error) {
  var locations []Location
  if err := db.Cypher(&neoism.CypherQuery{
    Statement: `MATCH (location:Location)
                RETURN DISTINCT ID(location) as id, location.name as name`,
    Result: &locations,
  }); err != nil {
    return nil, err
  }
  return &locations, nil
}
