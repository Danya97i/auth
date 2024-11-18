package user

import (
	"context"

	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Masterminds/squirrel"
)

// Delete удаляет пользователя по id
func (r *repo) Delete(ctx context.Context, id int64) error {
	deleteUserQueryBuilder := squirrel.Delete("users").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": id})

	deleteUserQuery, args, err := deleteUserQueryBuilder.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{RawQuery: deleteUserQuery}
	if _, err := r.db.DB().ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
