package user

import (
	"context"
	"time"

	"github.com/Danya97i/auth/internal/client/db"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/repository"
	"github.com/Danya97i/auth/internal/repository/user/converter"
	repoModels "github.com/Danya97i/auth/internal/repository/user/models"
	"github.com/Masterminds/squirrel"
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

// Create создает нового пользователя
func (r *repo) Create(ctx context.Context, userInfo models.UserInfo, passHash string) (int64, error) {
	insertUserQueryBuilder := squirrel.Insert("users").
		PlaceholderFormat(squirrel.Dollar).
		Columns("name", "email", "password", "role").
		Values(userInfo.Name, userInfo.Email, passHash, userInfo.Role.String()).Suffix("RETURNING id")
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
	query := db.Query{RawQuery: getUserQuery}
	var user repoModels.User
	if err := r.db.DB().ScanOneContext(ctx, &user, query, args...); err != nil {
		return nil, err
	}
	return converter.ToUserFromRepo(user), nil
}

// Update обновляет данные пользователя
func (r *repo) Update(ctx context.Context, id int64, userInfo models.UserInfo) error {
	updateUserQueryBuilder := squirrel.Update("users").
		PlaceholderFormat(squirrel.Dollar).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": id})
	if userInfo.Name != "" {
		updateUserQueryBuilder = updateUserQueryBuilder.Set("name", userInfo.Name)
	}
	if userInfo.Role != models.UNKNOWN {
		updateUserQueryBuilder = updateUserQueryBuilder.Set("role", userInfo.Role.String())
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
