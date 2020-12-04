package pgsql

import (
	"github.com/go-pg/pg/v9/orm"

	"github.com/blueskyinterfaces/secundusapi"
)

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u User) View(db orm.DB, id int) (secundusapi.User, error) {
	user := secundusapi.User{Base: secundusapi.Base{ID: id}}
	err := db.Select(&user)
	return user, err
}

// Update updates user's info
func (u User) Update(db orm.DB, user secundusapi.User) error {
	return db.Update(&user)
}
