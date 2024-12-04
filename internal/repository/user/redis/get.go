package redis

import (
	"context"
	"fmt"
	"strconv"

	redigo "github.com/gomodule/redigo/redis"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/repository/user/redis/converter"
	repoModels "github.com/Danya97i/auth/internal/repository/user/redis/models"
)

// Get implements repository.UserCache.
func (r *repo) Get(ctx context.Context, id int64) (*models.User, error) {
	idStr := strconv.FormatInt(id, 10)
	values, err := r.cli.HGetAll(ctx, idStr)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, fmt.Errorf("user not found") //model.ErrorNoteNotFound
	}

	var user repoModels.User
	err = redigo.ScanStruct(values, &user)
	if err != nil {
		return nil, err
	}
	return converter.ToUserFromRepo(user), nil
}
