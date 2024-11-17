package user

import (
	"github.com/Danya97i/auth/internal/client/db"
	"github.com/Danya97i/auth/internal/repository"
)

type repo struct {
	db db.Client
}

// NewRepository создает новый user repository
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
	}
}
