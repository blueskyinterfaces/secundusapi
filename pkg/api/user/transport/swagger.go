package transport

import (
	"github.com/blueskyinterfaces/secundusapi"
)

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*secundusapi.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []secundusapi.User `json:"users"`
		Page  int                `json:"page"`
	}
}
