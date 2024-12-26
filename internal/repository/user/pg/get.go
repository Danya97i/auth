package user

import (
	"context"

	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Masterminds/squirrel"

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/repository/user/pg/converter"
	repoModels "github.com/Danya97i/auth/internal/repository/user/pg/models"
)

// User возвращает пользователя по id
func (r *repo) User(ctx context.Context, id int64) (*models.User, error) {
	getUserQueryBuilder := squirrel.Select("id", "name", "email", "role", "created_at", "updated_at").
		PlaceholderFormat(squirrel.Dollar).
		From("users").
		Where(squirrel.Eq{"id": id})

	getUserQuery, args, err := getUserQueryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var user repoModels.User
	query := db.Query{RawQuery: getUserQuery}
	if err := r.db.DB().ScanOneContext(ctx, &user, query, args...); err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) UserByName(ctx context.Context, name string) (*models.User, error) {
	getUserQueryBuilder := squirrel.Select("id", "name", "email", "role", "created_at", "updated_at", "password").
		PlaceholderFormat(squirrel.Dollar).
		From("users").
		Where(squirrel.Eq{"name": name})

	getUserQuery, args, err := getUserQueryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var user repoModels.User
	query := db.Query{RawQuery: getUserQuery}
	if err := r.db.DB().ScanOneContext(ctx, &user, query, args...); err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
