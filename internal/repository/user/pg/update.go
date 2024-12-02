package user

import (
	"context"
	"time"

	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Masterminds/squirrel"

	"github.com/Danya97i/auth/internal/models"
)

// Update обновляет данные пользователя
func (r *repo) Update(ctx context.Context, id int64, userInfo models.UserInfo) error {
	updateUserQueryBuilder := squirrel.Update("users").
		PlaceholderFormat(squirrel.Dollar).
		Set("role", userInfo.Role).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": id})

	if userInfo.Name != nil {
		updateUserQueryBuilder = updateUserQueryBuilder.Set("name", userInfo.Name)
	}

	updateUserQuery, args, err := updateUserQueryBuilder.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{RawQuery: updateUserQuery}
	if _, err := r.db.DB().ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
