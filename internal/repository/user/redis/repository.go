package redis

import (
	"github.com/Danya97i/auth/internal/client/cache"
	"github.com/Danya97i/auth/internal/repository"
)

type repo struct {
	cli cache.RedisClient
}

// NewRepositoty creates a new repository
func NewRepositoty(cli cache.RedisClient) repository.UserCache {
	return &repo{cli: cli}

}
