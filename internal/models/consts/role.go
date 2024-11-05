package consts

// Role – роль пользователя
type Role string

// Список ролей
const (
	Unknown Role = "unknown"
	Admin   Role = "admin"
	User    Role = "user"
)
