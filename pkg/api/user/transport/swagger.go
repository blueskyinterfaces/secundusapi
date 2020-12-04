package transport

import (
	"github.com/secundusteam/secundus"
)

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*secundus.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []secundus.User `json:"users"`
		Page  int             `json:"page"`
	}
}
