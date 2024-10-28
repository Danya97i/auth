package logs

import (
	"context"

	"github.com/Masterminds/squirrel"

	"github.com/Danya97i/auth/internal/client/db"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.LogRepository {
	return &repo{db: db}
}

func (r *repo) Save(ctx context.Context, logInfo models.LogInfo) error {
	insertLogQueryBuilder := squirrel.Insert("users_logs").
		PlaceholderFormat(squirrel.Dollar).
		Columns("action", "user_id").
		Values(logInfo.Action, logInfo.UserID)

	insertLogQuery, args, err := insertLogQueryBuilder.ToSql()
	if err != nil {
		return err
	}
	query := db.Query{RawQuery: insertLogQuery}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}