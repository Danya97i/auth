package user

import (
	"context"
	"fmt"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
)

// UpdateUser обновляет пользователя по id
func (s *service) UpdateUser(ctx context.Context, id int64, userInfo *models.UserInfo) error {
	if userInfo == nil {
		return fmt.Errorf("userInfo is empty")
	}
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var txErr error
		txErr = s.userRepo.Update(ctx, id, *userInfo)
		if txErr != nil {
			return txErr
		}

		// создаем лог
		txErr = s.logRepo.Save(ctx, models.LogInfo{
			UserID: id,
			Action: consts.ActionUpdate,
		})

		return txErr
	})
	return err
}
