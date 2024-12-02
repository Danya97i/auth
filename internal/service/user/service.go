package user

import (
	"github.com/Danya97i/platform_common/pkg/db"

	"github.com/Danya97i/auth/internal/repository"
	serv "github.com/Danya97i/auth/internal/service"
)

type service struct {
	userRepo  repository.UserRepository
	logRepo   repository.LogRepository
	txManager db.TxManager
	userCache repository.UserCache
}

// NewService создает новый user service
func NewService(
	userRepo repository.UserRepository,
	logRepo repository.LogRepository,
	txManager db.TxManager,
	userCache repository.UserCache,
) serv.UserService {
	return &service{
		userRepo:  userRepo,
		logRepo:   logRepo,
		txManager: txManager,
		userCache: userCache,
	}
}
