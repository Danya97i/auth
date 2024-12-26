package auth

import (
	"time"

	"github.com/Danya97i/auth/internal/repository"
)

type service struct {
	userRepo repository.UserRepository

	accessTokenSecret  string
	refreshTokenSecret string

	accessTokenExpiration  time.Duration
	refreshTokenExpiration time.Duration
}

// NewService creates a new service
func NewService(
	userRepo repository.UserRepository,
	accessTokenSecret string,
	refreshTokenSecret string,

	accessTokenExpiration time.Duration,
	refreshTokenExpiration time.Duration,
) *service {
	return &service{
		userRepo:               userRepo,
		accessTokenSecret:      accessTokenSecret,
		refreshTokenSecret:     refreshTokenSecret,
		accessTokenExpiration:  accessTokenExpiration,
		refreshTokenExpiration: refreshTokenExpiration,
	}
}
