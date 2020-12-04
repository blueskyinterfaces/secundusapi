package mock

import (
	"github.com/blueskyinterfaces/secundusapi"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(secundusapi.User) (string, error)
}

// GenerateToken mock
func (j JWT) GenerateToken(u secundusapi.User) (string, error) {
	return j.GenerateTokenFn(u)
}
