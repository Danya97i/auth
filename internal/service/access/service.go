package access

import (
	"github.com/Danya97i/auth/internal/repository"
)

type service struct {
	accessRuleRepo    repository.AccessRuleRepository
	accessTokenSecret string
}

// NewService creates a new service
func NewService(
	accessRuleRepo repository.AccessRuleRepository,
	accessTokenSecret string,
) *service {
	return &service{
		accessRuleRepo:    accessRuleRepo,
		accessTokenSecret: accessTokenSecret,
	}
}
