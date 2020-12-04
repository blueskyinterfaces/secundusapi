package mockdb

import (
	"github.com/go-pg/pg/v9/orm"

	"github.com/secundusteam/secundus"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, secundus.User) (secundus.User, error)
	ViewFn           func(orm.DB, int) (secundus.User, error)
	FindByUsernameFn func(orm.DB, string) (secundus.User, error)
	FindByTokenFn    func(orm.DB, string) (secundus.User, error)
	ListFn           func(orm.DB, *secundus.ListQuery, secundus.Pagination) ([]secundus.User, error)
	DeleteFn         func(orm.DB, secundus.User) error
	UpdateFn         func(orm.DB, secundus.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr secundus.User) (secundus.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (secundus.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (secundus.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (secundus.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *secundus.ListQuery, p secundus.Pagination) ([]secundus.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr secundus.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr secundus.User) error {
	return u.UpdateFn(db, usr)
}
