package models

// User model
type User struct {
	ID          int64  `redis:"id"`
	Name        string `redis:"name"`
	Role        string `redis:"role"`
	Email       string `redis:"email"`
	CreatedAtNs int64  `redis:"created_at"`
	UpdatedAtNS *int64 `redis:"updated_at"`
}
