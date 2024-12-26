package auth

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/Danya97i/auth/internal/utils"
)

func (s *service) Login(ctx context.Context, username, password string) (string, error) {
	u, err := s.userRepo.UserByName(ctx, username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PassHash), []byte(password)); err != nil {
		return "", err
	}

	refreshToken, err := utils.GenerateToken(*u.Info, []byte(s.refreshTokenSecret), s.refreshTokenExpiration)

	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return refreshToken, nil
}
