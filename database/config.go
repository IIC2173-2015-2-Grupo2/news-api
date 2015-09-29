package database

import (
	"log"

	"github.com/jmcvetta/neoism"
)

/*
Connect to database
*/
func Connect(user, password, host, port string) *neoism.Database {
	uri := "http://" + user + ":" + password + "@" + host + ":" + port + "/db/data"
	db, err := neoism.Connect(uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
