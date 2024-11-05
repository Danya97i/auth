package user

import (
	"context"
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
)

// CreateUser создает нового пользователя
func (s *service) CreateUser(ctx context.Context, userInfo models.UserInfo, pass string, passConfirm string) (int64, error) {
	if userInfo.Name == nil {
		return 0, errors.New("user name is empty")
	}

	_, err := mail.ParseAddress(userInfo.Email)
	if err != nil {
		return 0, err
	}

	if pass != passConfirm {
		return 0, errors.New("passwords don't match")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	var id int64
	err = s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var txErr error
		id, txErr = s.userRepo.Create(ctx, userInfo, string(passHash))
		if err != nil {
			return txErr
		}

		// создаем лог
		txErr = s.logRepo.Save(ctx, models.LogInfo{
			UserID: id,
			Action: consts.ActionCreate,
		})
		if txErr != nil {
			return txErr
		}

		return nil
	})

	return id, err
}
