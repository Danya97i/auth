package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Danya97i/auth/internal/client/db"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/repository"
	serv "github.com/Danya97i/auth/internal/service"
)

type service struct {
	userRepo  repository.UserRepository
	logRepo   repository.LogRepository
	txManager db.TxManager
}

// NewService создает новый user service
func NewService(
	userRepo repository.UserRepository,
	logRepo repository.LogRepository,
	txManager db.TxManager,
) serv.UserService {
	return &service{
		userRepo:  userRepo,
		logRepo:   logRepo,
		txManager: txManager,
	}
}

// CreateUser создает нового пользователя
func (s *service) CreateUser(ctx context.Context, userInfo models.UserInfo, pass string, passConfirm string) (int64, error) {
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
			Action: models.ActionCreate,
		})
		if txErr != nil {
			return txErr
		}
		return nil
	})

	return id, err
}

// User возвращает пользователя по id
func (s *service) User(ctx context.Context, id int64) (*models.User, error) {
	var user *models.User
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var txErr error
		user, txErr = s.userRepo.User(ctx, id)
		if txErr != nil {
			return txErr
		}

		// создаем лог
		txErr = s.logRepo.Save(ctx, models.LogInfo{
			UserID: id,
			Action: models.ActionGet,
		})
		return txErr
	})
	return user, err
}

// UpdateUser обновляет пользователя по id
func (s *service) UpdateUser(ctx context.Context, id int64, userInfo *models.UserInfo) error {
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var txErr error
		txErr = s.userRepo.Update(ctx, id, *userInfo)
		if txErr != nil {
			return txErr
		}

		// создаем лог
		txErr = s.logRepo.Save(ctx, models.LogInfo{
			UserID: id,
			Action: models.ActionUpdate,
		})
		return txErr
	})
	return err
}

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
			Action: models.ActionDelete,
		})
		return txErr
	})
	return err
}
