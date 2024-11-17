package models

import (
	"time"

	"github.com/Danya97i/auth/internal/models/consts"
)

// User – пользователь
type User struct {
	ID        int64
	Info      *UserInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// UserInfo – информация о пользователе
type UserInfo struct {
	Name  *string
	Email string
	Role  consts.Role
}

// Role – роль пользователя
type Role int32

// Roles список ролей
const (
	UNKNOWN Role = iota
	ADMIN
	USER
)

func (r Role) String() string {
	switch r {
	case ADMIN:
		return "admin"
	case USER:
		return "user"
	default:
		return "unknown"
	}
}
