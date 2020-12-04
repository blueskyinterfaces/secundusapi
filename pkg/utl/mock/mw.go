package mock

import (
	"github.com/secundusteam/secundus"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(secundus.User) (string, error)
}

// GenerateToken mock
func (j JWT) GenerateToken(u secundus.User) (string, error) {
	return j.GenerateTokenFn(u)
}
