package access_rule

import (
	"github.com/Danya97i/platform_common/pkg/db"

	"github.com/Danya97i/auth/internal/repository"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new repository
func NewRepository(db db.Client) repository.AccessRuleRepository {
	return &repo{
		db: db,
	}
}
