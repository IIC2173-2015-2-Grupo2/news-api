package models

import (
	"github.com/jmcvetta/neoism"
)

/*
Company model
*/
type Company struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ---------------------------------------------------------------------------

/*
GetCompanies returns collection of news
*/
func GetCompanies(db *neoism.Database) (*[]Company, error) {
	var companies []Company
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (company:Company)
                RETURN ID(company) as id, company.name as name`,
		Result: &companies,
	}); err != nil {
		return nil, err
	}
	return &companies, nil
}
