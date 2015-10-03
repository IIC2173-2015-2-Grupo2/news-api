package models

import (
	"errors"

	"github.com/jmcvetta/neoism"
)

/*
User model
*/
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ---------------------------------------------------------------------------

/*
Create user on database
*/
func (u *User) Create(db *neoism.Database) error {
	return nil
}

/*
GetUser returns the user with that id
*/
func GetUser(db *neoism.Database, id int) (*User, error) {
	var users []User
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (user:User)
								WHERE ID(user) = {id}
								RETURN user.name as name, user.email as email`,
		Parameters: neoism.Props{"id": id},
		Result:     &users,
	}); err != nil {
		return nil, err

	} else if len(users) == 0 {
		return nil, errors.New("not found")

	} else {
		return &users[0], nil
	}
}

/*
GetUsers returns collection of users
*/
func GetUsers(db *neoism.Database) (*[]User, error) {
	var users []User
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (new:User)
								RETURN user.name as name, user.username as username, user.email as email`,
		Result: &users,
	}); err != nil {
		return nil, err
	}
	return &users, nil
}

/*
FindUserByUsername find user
*/
func FindUserByUsername(db *neoism.Database, username string) (*User, error) {
	var users []User
	if err := db.Cypher(&neoism.CypherQuery{
		Statement: `MATCH (user:User)
								WHERE user.username = {username}
								RETURN user.name as name, user.username as username, user.email as email, user.password as password`,
		Parameters: neoism.Props{"username": username},
		Result:     &users,
	}); err != nil {
		return nil, err

	} else if len(users) == 0 {
		return nil, errors.New("not found")

	} else {
		return &users[0], nil
	}
}
