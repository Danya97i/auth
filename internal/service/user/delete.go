package user

import (
	"context"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
)

// DeleteUser удаляет пользователя по id
func (s *service) DeleteUser(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var txErr error
		txErr = s.userRepo.Delete(ctx, id)
		if txErr != nil {
			return txErr
		}

		// создаем лог
		txErr = s.logRepo.Save(ctx, models.LogInfo{
			UserID: id,
			Action: consts.ActionDelete,
		})
		return txErr
	})
	return err
}
