package auth

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/secundusteam/secundus"
)

// Custom errors
var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Username or password does not exist")
)

// Authenticate tries to authenticate the user provided by username and password
func (a Auth) Authenticate(c echo.Context, user, pass string) (secundus.AuthToken, error) {
	u, err := a.udb.FindByUsername(a.db, user)
	if err != nil {
		return secundus.AuthToken{}, err
	}

	if !a.sec.HashMatchesPassword(u.Password, pass) {
		return secundus.AuthToken{}, ErrInvalidCredentials
	}

	if !u.Active {
		return secundus.AuthToken{}, secundus.ErrUnauthorized
	}

	token, err := a.tg.GenerateToken(u)
	if err != nil {
		return secundus.AuthToken{}, secundus.ErrUnauthorized
	}

	u.UpdateLastLogin(a.sec.Token(token))

	if err := a.udb.Update(a.db, u); err != nil {
		return secundus.AuthToken{}, err
	}

	return secundus.AuthToken{Token: token, RefreshToken: u.Token}, nil
}

// Refresh refreshes jwt token and puts new claims inside
func (a Auth) Refresh(c echo.Context, refreshToken string) (string, error) {
	user, err := a.udb.FindByToken(a.db, refreshToken)
	if err != nil {
		return "", err
	}
	return a.tg.GenerateToken(user)
}

// Me returns info about currently logged user
func (a Auth) Me(c echo.Context) (secundus.User, error) {
	au := a.rbac.User(c)
	return a.udb.View(a.db, au.ID)
}
