package query

import (
	"github.com/labstack/echo"

	"github.com/secundusteam/secundus"
)

// List prepares data for list queries
func List(u secundus.AuthUser) (*secundus.ListQuery, error) {
	switch true {
	case u.Role <= secundus.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == secundus.CompanyAdminRole:
		return &secundus.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == secundus.LocationAdminRole:
		return &secundus.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
