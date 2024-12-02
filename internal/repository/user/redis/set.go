package redis

import (
	"context"
	"strconv"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/repository/user/redis/converter"
)

// Set sets user in redis
func (r *repo) Set(ctx context.Context, user *models.User) error {
	repoUser := converter.ToRepoUserFromService(user)
	if repoUser == nil {
		return nil
	}

	key := strconv.FormatInt(user.ID, 10)

	err := r.cli.HashSet(ctx, key, repoUser)
	if err != nil {
		return nil
	}

	return nil
}
