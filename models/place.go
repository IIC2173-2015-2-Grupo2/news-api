package models

import (
	"github.com/jmcvetta/neoism"
)

/*
Place model
*/
type Place struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ---------------------------------------------------------------------------

/*
GetPlaces returns collection of news
*/
func GetPlaces(db *neoism.Database) (*[]Place, error) {
	var places []Place
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (place:Place)
                RETURN ID(place) as id, place.name as name`,
		Result: &places,
	}); err != nil {
		return nil, err
	}
	return &places, nil
}
