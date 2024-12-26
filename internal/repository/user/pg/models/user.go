package models

import (
	"database/sql"
)

// User – тип пользователя для хранения в БД
type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	Password  string       `db:"password"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// UserInfo – информация о пользователе для хранения в БД
type UserInfo struct {
	Name  *string `db:"name"`
	Email string  `db:"email"`
	Role  string  `db:"role"`
}
