package controllers

import (
	"fmt"

	"github.com/jmcvetta/neoism"
	"github.com/jpillora/go-ogle-analytics"
)

/*
Base CRUD
*/
type Base struct {
	DB        *neoism.Database
	Analytics *ga.Client
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
