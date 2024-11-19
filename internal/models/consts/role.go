package consts

// Role – роль пользователя
type Role string

// Список ролей
const (
	Unknown Role = "UNKNOWN"
	Admin   Role = "ADMIN"
	User    Role = "USER"
)
