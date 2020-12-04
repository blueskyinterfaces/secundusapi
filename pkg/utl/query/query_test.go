package query_test

import (
	"testing"

	"github.com/labstack/echo"

	"github.com/secundusteam/secundus"

	"github.com/stretchr/testify/assert"

	"github.com/secundusteam/secundus/pkg/utl/query"
)

func TestList(t *testing.T) {
	type args struct {
		user secundus.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *secundus.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: secundus.AuthUser{
				Role: secundus.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: secundus.AuthUser{
				Role:      secundus.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &secundus.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: secundus.AuthUser{
				Role:       secundus.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &secundus.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: secundus.AuthUser{
				Role: secundus.UserRole,
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
