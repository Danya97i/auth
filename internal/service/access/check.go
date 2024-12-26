package access

import (
	"context"
	"fmt"

	"github.com/Danya97i/auth/internal/models/consts"
	"github.com/Danya97i/auth/internal/utils"
)

// Check checks the access token
func (s service) Check(ctx context.Context, accessToken string, endpoint string) error {
	claims, err := utils.VerifyToken(accessToken, []byte(s.accessTokenSecret))
	if err != nil {
		return err
	}

	// идем в базу и достаем для роли из токена список доступных ендпоинтов
	if err := s.accessRuleRepo.CheckRuleExist(ctx, consts.Role(claims.Role), endpoint); err != nil {
		return fmt.Errorf("access denied: %w", err)
	}

	return nil
}
