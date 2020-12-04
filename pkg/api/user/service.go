package user

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	"github.com/secundusteam/secundus"
	"github.com/secundusteam/secundus/pkg/api/user/platform/pgsql"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, secundus.User) (secundus.User, error)
	List(echo.Context, secundus.Pagination) ([]secundus.User, error)
	View(echo.Context, int) (secundus.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, Update) (secundus.User, error)
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
	Create(orm.DB, secundus.User) (secundus.User, error)
	View(orm.DB, int) (secundus.User, error)
	List(orm.DB, *secundus.ListQuery, secundus.Pagination) ([]secundus.User, error)
	Update(orm.DB, secundus.User) error
	Delete(orm.DB, secundus.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) secundus.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, secundus.AccessRole, int, int) error
	IsLowerRole(echo.Context, secundus.AccessRole) error
}
