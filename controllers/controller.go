package controllers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jmcvetta/neoism"
	"github.com/jpillora/go-ogle-analytics"
)

/*
Base CRUD
*/
type Base struct {
	Analytics *ga.Client
}

/*
Neo4jBase Database base
*/
type Neo4jBase struct {
	Base
	DB *neoism.Database
}

/*
PgBase Database base
*/
type PgBase struct {
	Base
	DB *gorm.DB
}

/*
Log to Analytics server
*/
func (r Base) Log(category string, action string) {
	if r.Analytics == nil {
		// Do nothing
	} else if err := r.Analytics.Send(ga.NewEvent(category, action)); err != nil {
		fmt.Println("Could not log:", category, "-", action)
	}
}
