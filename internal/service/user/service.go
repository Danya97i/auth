package user

import (
	"github.com/Danya97i/auth/internal/client/db"
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
