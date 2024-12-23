package models

import "github.com/dgrijalva/jwt-go"

// UserClaims is the claims of the user
type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
