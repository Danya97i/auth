package converter

import (
	"time"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
	repoModels "github.com/Danya97i/auth/internal/repository/user/redis/models"
)

// ToUserFromRepo converts repoUser to models.User
func ToUserFromRepo(repoUser repoModels.User) *models.User {
	var updatedAt *time.Time

	if repoUser.UpdatedAtNS != nil {
		t := time.Unix(0, *repoUser.UpdatedAtNS)
		updatedAt = &t
	}

	return &models.User{
		ID: repoUser.ID,
		Info: &models.UserInfo{
			Name:  &repoUser.Name,
			Email: repoUser.Email,
			Role:  consts.Role(repoUser.Role),
		},
		CreatedAt: time.Unix(0, repoUser.CreatedAtNs),
		UpdatedAt: updatedAt,
	}
}

// ToRepoUserFromService converts models.User to repoModels.User
func ToRepoUserFromService(user *models.User) *repoModels.User {
	if user == nil {
		return nil
	}

	var updatedAtNs *int64
	if user.UpdatedAt != nil {
		t := user.UpdatedAt.UnixNano()
		updatedAtNs = &t
	}

	return &repoModels.User{
		ID:          user.ID,
		Name:        *user.Info.Name,
		Email:       user.Info.Email,
		Role:        string(user.Info.Role),
		CreatedAtNs: user.CreatedAt.UnixNano(),
		UpdatedAtNS: updatedAtNs,
	}
}
