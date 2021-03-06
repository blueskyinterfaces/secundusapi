// secundus
//
// API Docs for secundus v1
//
// 	 Terms Of Service:  N/A
//     Schemes: http
//     Host: localhost:3000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer: []
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package api

import (
	"crypto/sha1"

	"github.com/labstack/echo"
	"github.com/blueskyinterfaces/secundusapi/pkg/utl/zlog"

	"github.com/blueskyinterfaces/secundusapi/pkg/api/auth"
	al "github.com/blueskyinterfaces/secundusapi/pkg/api/auth/logging"
	at "github.com/blueskyinterfaces/secundusapi/pkg/api/auth/transport"
	"github.com/blueskyinterfaces/secundusapi/pkg/api/password"
	pl "github.com/blueskyinterfaces/secundusapi/pkg/api/password/logging"
	pt "github.com/blueskyinterfaces/secundusapi/pkg/api/password/transport"
	"github.com/blueskyinterfaces/secundusapi/pkg/api/user"
	ul "github.com/blueskyinterfaces/secundusapi/pkg/api/user/logging"
	ut "github.com/blueskyinterfaces/secundusapi/pkg/api/user/transport"

	"github.com/blueskyinterfaces/secundusapi/pkg/utl/config"
	"github.com/blueskyinterfaces/secundusapi/pkg/utl/jwt"
	authMw "github.com/blueskyinterfaces/secundusapi/pkg/utl/middleware/auth"
	"github.com/blueskyinterfaces/secundusapi/pkg/utl/postgres"
	"github.com/blueskyinterfaces/secundusapi/pkg/utl/rbac"
	"github.com/blueskyinterfaces/secundusapi/pkg/utl/secure"
	"github.com/blueskyinterfaces/secundusapi/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	db, err := postgres.New("postgres://secunduswebuser:111fart@secundata.eastus2.cloudapp.azure.com:5432/secundusweb", cfg.DB.Timeout, cfg.DB.LogQueries) // TODO env it
	if err != nil {
		return err
	}

	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.Service{}
	jwt, err := jwt.New(cfg.JWT.SigningAlgorithm, "oncetherewasamannamedkeanureeveswhowasthegreteastactorofalltimeeventuallyhefiguredoutthesnwertotheuniversesgreatestquestionwhatthefuck?", cfg.JWT.DurationMinutes, cfg.JWT.MinSecretLength) // TODO env it
	if err != nil {
		return err
	}

	log := zlog.New()

	e := server.New()
	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	e.HTTPErrorHandler = customHTTPErrorHandler

	authMiddleware := authMw.Middleware(jwt)

	at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, authMiddleware)

	v1 := e.Group("/v1")
	v1.Use(authMiddleware)

	ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1)
	pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}

func customHTTPErrorHandler(err error, c echo.Context) {
	errorPage := "assets/app/dist/index.html"
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
