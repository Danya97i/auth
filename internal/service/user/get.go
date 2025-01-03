package user

import (
	"context"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
)

// User возвращает пользователя по id
func (s *service) User(ctx context.Context, id int64) (*models.User, error) {
	var user *models.User
	var err error

	// пытаемся получить запись из кэша
	user, err = s.userCache.Get(ctx, id)
	if err == nil {
		return user, nil
	}

	// если попали сюда, значит в кэше пользователя нет
	err = s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var txErr error
		user, txErr = s.userRepo.User(ctx, id)
		if txErr != nil {
			return txErr
		}

		// создаем лог
		txErr = s.logRepo.Save(ctx, models.LogInfo{
			UserID: id,
			Action: consts.ActionGet,
		})
		return txErr
	})
	if err != nil {
		return nil, err
	}

	// записываем в кэш
	_ = s.userCache.Set(ctx, user)

	return user, nil
}
