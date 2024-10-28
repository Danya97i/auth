package user

import (
	"context"
	"errors"

	"github.com/Danya97i/auth/internal/client/db"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/repository"
	serv "github.com/Danya97i/auth/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo  repository.UserRepository
	txManager db.TxManager
}

// NewService создает новый user service
func NewService(userRepo repository.UserRepository, txManager db.TxManager) serv.UserService {
	return &service{
		userRepo:  userRepo,
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
	id, err := s.userRepo.Create(ctx, userInfo, string(passHash))
	if err != nil {
		return 0, err
	}
	return id, nil
}

// User возвращает пользователя по id
func (s *service) User(ctx context.Context, id int64) (*models.User, error) {
	return s.userRepo.User(ctx, id)
}

// UpdateUser обновляет пользователя по id
func (s *service) UpdateUser(ctx context.Context, id int64, userInfo *models.UserInfo) error {
	return s.userRepo.Update(ctx, id, *userInfo)
}

// DeleteUser удаляет пользователя по id
func (s *service) DeleteUser(ctx context.Context, id int64) error {
	return s.userRepo.Delete(ctx, id)
}
