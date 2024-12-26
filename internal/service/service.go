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

// AuthService интерфейс для работы с авторизацией
type AuthService interface {
	Login(ctx context.Context, username string, password string) (string, error)
	GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}

// AccessService интерфейс для работы с аутентификацией
type AccessService interface {
	Check(ctx context.Context, accessToken string, endpoint string) error
}
