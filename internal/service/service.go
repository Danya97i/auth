package service

import (
	"context"

	"github.com/Danya97i/auth/internal/models"
)

// UserService интерфейс для работы с пользователями
type UserService interface {
	CreateUser(ctx context.Context, userInfo models.UserInfo, pass string, passConfirm string) (int64, error)
	User(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, id int64, userInfo *models.UserInfo) error
	DeleteUser(ctx context.Context, id int64) error
}
