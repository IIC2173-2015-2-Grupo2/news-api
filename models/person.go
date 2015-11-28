package models

import (
	"github.com/jmcvetta/neoism"
)

/*
Person model
*/
type Person struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ---------------------------------------------------------------------------

/*
GetPeople returns collection of news
*/
func GetPeople(db *neoism.Database) (*[]Person, error) {
	var people []Person
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (person:Person)
                RETURN ID(person) as id, person.name as name`,
		Result: &people,
	}); err != nil {
		return nil, err
	}
	return &people, nil
}
