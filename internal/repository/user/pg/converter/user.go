package converter

import (
	"strings"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
	repoModels "github.com/Danya97i/auth/internal/repository/user/pg/models"
)

// ToUserFromRepo конвертирует модель пользователя из репозитория в модель пользователя
func ToUserFromRepo(repoUser *repoModels.User) *models.User {
	if repoUser == nil {
		return nil
	}

	info := ToUserInfoFromRepo(repoUser.Info)

	user := &models.User{
		ID:        repoUser.ID,
		Info:      info,
		PassHash:  repoUser.Password,
		CreatedAt: repoUser.CreatedAt.Time,
	}

	if repoUser.UpdatedAt.Valid {
		user.UpdatedAt = &repoUser.UpdatedAt.Time
	}

	return user
}

// ToUserInfoFromRepo конвертирует информацию о пользователе из репозитория в модель пользователя
func ToUserInfoFromRepo(repoUserInfo repoModels.UserInfo) *models.UserInfo {
	userInfo := models.UserInfo{
		Email: repoUserInfo.Email,
		Role:  ToUserRoleFromRepo(repoUserInfo.Role),
	}

	if repoUserInfo.Name != nil {
		userInfo.Name = repoUserInfo.Name
	}

	return &userInfo
}

// ToUserRoleFromRepo конвертирует роль пользователя из репозитория в модель пользователя
func ToUserRoleFromRepo(repoUserRole string) consts.Role {
	return consts.Role(strings.ToUpper(repoUserRole))
}
