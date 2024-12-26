package auth

import (
	"context"

	"github.com/Danya97i/auth/internal/utils"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(s.refreshTokenSecret))
	if err != nil {
		return "", err
	}

	user, err := s.userRepo.UserByName(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateToken(*user.Info, []byte(s.accessTokenSecret), s.accessTokenExpiration)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
