package models

import "github.com/jinzhu/gorm"

/*
User model
*/
type User struct {
	ID       int    `gorm:"primary_key"`
	Name     string `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
	Username string `sql:"size:255"`
	Email    string `sql:"size:255"`
	Password string `sql:"size:255"`
}

// ---------------------------------------------------------------------------

/*
Save user on database
*/
func (u *User) Save(db *gorm.DB) (*User, error) {
	db.Create(&u)
	return u, nil
}

/*
GetUser returns the user with that id
*/
func GetUser(db *gorm.DB, id int) (*User, error) {
	var user User
	db.Where("ID = ?", id).First(&user)
	return &user, nil
}

/*
GetUsers returns collection of users
*/
func GetUsers(db *gorm.DB) (*[]User, error) {
	var users []User
	db.Find(&users)
	return &users, nil
}

/*
FindUserByUsername find user
*/
func FindUserByUsername(db *gorm.DB, username string) (*User, error) {
	var users []User
	db.Where("username = ?", username).Find(&users)
	if len(users) > 0 {
		return &users[0], nil
	}
	return nil, nil
}
