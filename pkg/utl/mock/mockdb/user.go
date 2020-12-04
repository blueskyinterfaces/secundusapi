package mockdb

import (
	"github.com/go-pg/pg/v9/orm"

	"github.com/blueskyinterfaces/secundusapi"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, secundusapi.User) (secundusapi.User, error)
	ViewFn           func(orm.DB, int) (secundusapi.User, error)
	FindByUsernameFn func(orm.DB, string) (secundusapi.User, error)
	FindByTokenFn    func(orm.DB, string) (secundusapi.User, error)
	ListFn           func(orm.DB, *secundusapi.ListQuery, secundusapi.Pagination) ([]secundusapi.User, error)
	DeleteFn         func(orm.DB, secundusapi.User) error
	UpdateFn         func(orm.DB, secundusapi.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr secundusapi.User) (secundusapi.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (secundusapi.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (secundusapi.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (secundusapi.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *secundusapi.ListQuery, p secundusapi.Pagination) ([]secundusapi.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr secundusapi.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr secundusapi.User) error {
	return u.UpdateFn(db, usr)
}
