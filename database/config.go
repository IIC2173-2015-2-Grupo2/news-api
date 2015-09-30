package database

import (
	"fmt"

	"github.com/jmcvetta/neoism"
)

/*
Connect to database
*/
func Connect(user, password, host, port string) (*neoism.Database, error) {
	uri := "http://" + user + ":" + password + "@" + host + ":" + port + "/db/data"
	fmt.Println(uri)
	return neoism.Connect(uri)
}
