package query_test

import (
	"testing"

	"github.com/labstack/echo"

	"github.com/blueskyinterfaces/secundusapi"

	"github.com/stretchr/testify/assert"

	"github.com/blueskyinterfaces/secundusapi/pkg/utl/query"
)

func TestList(t *testing.T) {
	type args struct {
		user secundusapi.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *secundusapi.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: secundusapi.AuthUser{
				Role: secundusapi.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: secundusapi.AuthUser{
				Role:      secundusapi.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &secundusapi.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: secundusapi.AuthUser{
				Role:       secundusapi.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &secundusapi.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: secundusapi.AuthUser{
				Role: secundusapi.UserRole,
			}},
			wantErr: echo.ErrForbidden,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q, err := query.List(tt.args.user)
			assert.Equal(t, tt.wantData, q)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
