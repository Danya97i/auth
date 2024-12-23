package auth

import (
	"context"

	"github.com/Danya97i/auth/internal/utils"
)

func (s *service) GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
	claims, err := utils.VerifyToken(oldRefreshToken, []byte(s.refreshTokenSecret))
	if err != nil {
		return "", err
	}

	user, err := s.userRepo.UserByName(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	refreshToken, err := utils.GenerateToken(*user.Info, []byte(s.refreshTokenSecret), s.refreshTokenExpiration)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
