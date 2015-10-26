package models

import (
	"github.com/jmcvetta/neoism"
)

/*
NewsProvider model
*/
type NewsProvider struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ---------------------------------------------------------------------------

/*
GetNewsProviders returns collection of news
*/
func GetNewsProviders(db *neoism.Database) (*[]NewsProvider, error) {
	var newsproviders []NewsProvider
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (newsprovider:NewsProvider)
                RETURN ID(newsprovider) as id, newsprovider.name as name`,
		Result: &newsproviders,
	}); err != nil {
		return nil, err
	}
	return &newsproviders, nil
}
