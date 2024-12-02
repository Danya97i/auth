package user

import (
	"context"
	"database/sql"
	"strings"

	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Masterminds/squirrel"

	"github.com/Danya97i/auth/internal/models"
)

// Create создает нового пользователя
func (r *repo) Create(ctx context.Context, userInfo models.UserInfo, passHash string) (int64, error) {
	var userName sql.NullString

	if userInfo.Name != nil {
		userName = sql.NullString{
			String: *userInfo.Name,
			Valid:  true,
		}
	}

	insertUserQueryBuilder := squirrel.Insert("users").
		PlaceholderFormat(squirrel.Dollar).
		Columns("name", "email", "password", "role").
		Values(userName, userInfo.Email, passHash, strings.ToLower(string(userInfo.Role))).
		Suffix("RETURNING id")

	insertUserQuery, args, err := insertUserQueryBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	query := db.Query{RawQuery: insertUserQuery}
	var id int64
	if err := r.db.DB().ScanOneContext(ctx, &id, query, args...); err != nil {
		return 0, err
	}

	return id, nil
}
