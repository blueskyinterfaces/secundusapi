package user

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	"github.com/blueskyinterfaces/secundusapi"
	"github.com/blueskyinterfaces/secundusapi/pkg/api/user/platform/pgsql"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, secundusapi.User) (secundusapi.User, error)
	List(echo.Context, secundusapi.Pagination) ([]secundusapi.User, error)
	View(echo.Context, int) (secundusapi.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, Update) (secundusapi.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.User{}, rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, secundusapi.User) (secundusapi.User, error)
	View(orm.DB, int) (secundusapi.User, error)
	List(orm.DB, *secundusapi.ListQuery, secundusapi.Pagination) ([]secundusapi.User, error)
	Update(orm.DB, secundusapi.User) error
	Delete(orm.DB, secundusapi.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) secundusapi.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, secundusapi.AccessRole, int, int) error
	IsLowerRole(echo.Context, secundusapi.AccessRole) error
}
