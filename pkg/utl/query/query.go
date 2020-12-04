package query

import (
	"github.com/labstack/echo"

	"github.com/blueskyinterfaces/secundusapi"
)

// List prepares data for list queries
func List(u secundusapi.AuthUser) (*secundusapi.ListQuery, error) {
	switch true {
	case u.Role <= secundusapi.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == secundusapi.CompanyAdminRole:
		return &secundusapi.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == secundusapi.LocationAdminRole:
		return &secundusapi.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
