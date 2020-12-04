package pgsql_test

import (
	"testing"

	"github.com/secundusteam/secundus"
	"github.com/secundusteam/secundus/pkg/utl/mock"

	"github.com/secundusteam/secundus/pkg/api/auth/platform/pgsql"

	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	cases := []struct {
		name     string
		wantErr  bool
		id       int
		wantData secundus.User
	}{
		{
			name:    "User does not exist",
			wantErr: true,
			id:      1000,
		},
		{
			name: "Success",
			id:   2,
			wantData: secundus.User{
				Email:      "tomjones@mail.com",
				FirstName:  "Tom",
				LastName:   "Jones",
				Username:   "tomjones",
				RoleID:     1,
				CompanyID:  1,
				LocationID: 1,
				Password:   "newPass",
				Base: secundus.Base{
					ID: 2,
				},
				Role: &secundus.Role{
					ID:          1,
					AccessLevel: 1,
					Name:        "SUPER_ADMIN",
				},
			},
		},
	}

	dbCon := mock.NewPGContainer(t)
	defer dbCon.Shutdown()

	db := mock.NewDB(t, dbCon, &secundus.Role{}, &secundus.User{})

	if err := mock.InsertMultiple(db, &secundus.Role{
		ID:          1,
		AccessLevel: 1,
		Name:        "SUPER_ADMIN"}, &cases[1].wantData); err != nil {
		t.Error(err)
	}

	udb := pgsql.User{}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			user, err := udb.View(db, tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantData.ID != 0 {
				if user.ID == 0 {
					t.Errorf("empty response: %v", err)
				} else {
					tt.wantData.CreatedAt = user.CreatedAt
					tt.wantData.UpdatedAt = user.UpdatedAt
					assert.Equal(t, tt.wantData, user)
				}
			}
		})
	}
}

func TestFindByUsername(t *testing.T) {
	cases := []struct {
		name     string
		wantErr  bool
		username string
		wantData secundus.User
	}{
		{
			name:     "User does not exist",
			wantErr:  true,
			username: "notExists",
		},
		{
			name:     "Success",
			username: "tomjones",
			wantData: secundus.User{
				Email:      "tomjones@mail.com",
				FirstName:  "Tom",
				LastName:   "Jones",
				Username:   "tomjones",
				RoleID:     1,
				CompanyID:  1,
				LocationID: 1,
				Password:   "newPass",
				Base: secundus.Base{
					ID: 2,
				},
				Role: &secundus.Role{
					ID:          1,
					AccessLevel: 1,
					Name:        "SUPER_ADMIN",
				},
			},
		},
	}

	dbCon := mock.NewPGContainer(t)
	defer dbCon.Shutdown()

	db := mock.NewDB(t, dbCon, &secundus.Role{}, &secundus.User{})

	if err := mock.InsertMultiple(db, &secundus.Role{
		ID:          1,
		AccessLevel: 1,
		Name:        "SUPER_ADMIN"}, &cases[1].wantData); err != nil {
		t.Error(err)
	}

	udb := pgsql.User{}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			user, err := udb.FindByUsername(db, tt.username)
			assert.Equal(t, tt.wantErr, err != nil)

			if tt.wantData.ID != 0 {
				tt.wantData.CreatedAt = user.CreatedAt
				tt.wantData.UpdatedAt = user.UpdatedAt
				assert.Equal(t, tt.wantData, user)

			}
		})
	}
}

func TestFindByToken(t *testing.T) {
	cases := []struct {
		name     string
		wantErr  bool
		token    string
		wantData secundus.User
	}{
		{
			name:    "User does not exist",
			wantErr: true,
			token:   "notExists",
		},
		{
			name:  "Success",
			token: "loginrefresh",
			wantData: secundus.User{
				Email:      "johndoe@mail.com",
				FirstName:  "John",
				LastName:   "Doe",
				Username:   "johndoe",
				RoleID:     1,
				CompanyID:  1,
				LocationID: 1,
				Password:   "hunter2",
				Base: secundus.Base{
					ID: 1,
				},
				Role: &secundus.Role{
					ID:          1,
					AccessLevel: 1,
					Name:        "SUPER_ADMIN",
				},
				Token: "loginrefresh",
			},
		},
	}

	dbCon := mock.NewPGContainer(t)
	defer dbCon.Shutdown()

	db := mock.NewDB(t, dbCon, &secundus.Role{}, &secundus.User{})

	if err := mock.InsertMultiple(db, &secundus.Role{
		ID:          1,
		AccessLevel: 1,
		Name:        "SUPER_ADMIN"}, &cases[1].wantData); err != nil {
		t.Error(err)
	}

	udb := pgsql.User{}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			user, err := udb.FindByToken(db, tt.token)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantData.ID != 0 {
				tt.wantData.CreatedAt = user.CreatedAt
				tt.wantData.UpdatedAt = user.UpdatedAt
				assert.Equal(t, tt.wantData, user)

			}
		})
	}
}

func TestUpdate(t *testing.T) {
	cases := []struct {
		name     string
		wantErr  bool
		usr      secundus.User
		wantData secundus.User
	}{
		{
			name: "Success",
			usr: secundus.User{
				Base: secundus.Base{
					ID: 2,
				},
				FirstName: "Z",
				LastName:  "Freak",
				Address:   "Address",
				Phone:     "123456",
				Mobile:    "345678",
				Username:  "newUsername",
			},
			wantData: secundus.User{
				Email:      "tomjones@mail.com",
				FirstName:  "Z",
				LastName:   "Freak",
				Username:   "tomjones",
				RoleID:     1,
				CompanyID:  1,
				LocationID: 1,
				Password:   "newPass",
				Address:    "Address",
				Phone:      "123456",
				Mobile:     "345678",
				Base: secundus.Base{
					ID: 2,
				},
			},
		},
	}

	dbCon := mock.NewPGContainer(t)
	defer dbCon.Shutdown()

	db := mock.NewDB(t, dbCon, &secundus.Role{}, &secundus.User{})

	if err := mock.InsertMultiple(db, &secundus.Role{
		ID:          1,
		AccessLevel: 1,
		Name:        "SUPER_ADMIN"}, &cases[0].usr); err != nil {
		t.Error(err)
	}

	udb := pgsql.User{}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := udb.Update(db, tt.wantData)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantData.ID != 0 {
				user := secundus.User{
					Base: secundus.Base{
						ID: tt.usr.ID,
					},
				}
				if err := db.Select(&user); err != nil {
					t.Error(err)
				}
				tt.wantData.UpdatedAt = user.UpdatedAt
				tt.wantData.CreatedAt = user.CreatedAt
				tt.wantData.LastLogin = user.LastLogin
				tt.wantData.DeletedAt = user.DeletedAt
				assert.Equal(t, tt.wantData, user)
			}
		})
	}
}
