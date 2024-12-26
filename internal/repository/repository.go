package repository

import (
	"context"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
)

// UserRepository interface
type UserRepository interface {
	Create(ctx context.Context, userInfo models.UserInfo, passHash string) (int64, error)

	User(ctx context.Context, id int64) (*models.User, error)
	UserByName(ctx context.Context, name string) (*models.User, error)

	Update(ctx context.Context, id int64, user models.UserInfo) error
	Delete(ctx context.Context, id int64) error
}

// UserCache interface
type UserCache interface {
	Get(ctx context.Context, id int64) (*models.User, error)
	Set(ctx context.Context, user *models.User) error
}

// LogRepository interface
type LogRepository interface {
	Save(ctx context.Context, logInfo models.LogInfo) error
}

// AccessRuleRepository interface
type AccessRuleRepository interface {
	CheckRuleExist(ctx context.Context, role consts.Role, endpoint string) error
}
