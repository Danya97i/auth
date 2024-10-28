package converter

import (
	"github.com/Danya97i/auth/internal/models"
	repoModels "github.com/Danya97i/auth/internal/repository/user/models"
)

// ToUserFromRepo конвертирует модель пользователя из репозитория в модель пользователя
func ToUserFromRepo(repoUser repoModels.User) *models.User {
	user := &models.User{
		ID:        repoUser.ID,
		Info:      *ToUserInfoFromRepo(repoUser.Info),
		CreatedAt: repoUser.CreatedAt.Time,
	}
	if repoUser.UpdatedAt.Valid {
		user.UpdatedAt = &repoUser.UpdatedAt.Time
	}
	return user
}

// ToUserInfoFromRepo конвертирует информацию о пользователе из репозитория в модель пользователя
func ToUserInfoFromRepo(repoUserInfo repoModels.UserInfo) *models.UserInfo {
	return &models.UserInfo{
		Name:  repoUserInfo.Name,
		Email: repoUserInfo.Email,
		Role:  ToUserRoleFromRepo(repoUserInfo.Role),
	}
}

// ToUserRoleFromRepo конвертирует роль пользователя из репозитория в модель пользователя
func ToUserRoleFromRepo(repoUserRole string) models.Role {
	switch repoUserRole {
	case models.ADMIN.String():
		return models.ADMIN
	case models.USER.String():
		return models.USER
	default:
		return models.UNKNOWN
	}
}
