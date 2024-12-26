package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type accesTokenConfig struct {
	secret     string
	expiration int
}

type refreshTokenConfig struct {
	secret     string
	expiration int
}

// NewAccessTokenConfig returns a new access token config
func NewAccessTokenConfig() (*accesTokenConfig, error) {
	secret := os.Getenv("ACCESS_TOKEN_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("access token secret is empty")
	}

	expirationStr := os.Getenv("ACCESS_TOKEN_EXPIRATION")
	if expirationStr == "" {
		return nil, fmt.Errorf("access token expiration is empty")
	}

	expiration, err := strconv.Atoi(expirationStr)
	if err != nil {
		return nil, fmt.Errorf("access token expiration is not a number")
	}
	return &accesTokenConfig{secret, expiration}, nil
}

// NewRefreshTokenConfig returns a new refresh token config
func NewRefreshTokenConfig() (*refreshTokenConfig, error) {
	secret := os.Getenv("REFRESH_TOKEN_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("refresh token secret is empty")
	}

	expirationStr := os.Getenv("REFRESH_TOKEN_EXPIRATION")
	if expirationStr == "" {
		return nil, fmt.Errorf("refresh token expiration is empty")
	}

	expiration, err := strconv.Atoi(expirationStr)
	if err != nil {
		return nil, fmt.Errorf("refresh token expiration is not a number")
	}

	return &refreshTokenConfig{secret, expiration}, nil
}

// Secret returns the access token secret
func (c *accesTokenConfig) Secret() string {
	return c.secret
}

// Expiration returns the access token expiration
func (c *accesTokenConfig) Expiration() time.Duration {
	return time.Minute * time.Duration(c.expiration)
}

// Secret returns the refresh token secret
func (c *refreshTokenConfig) Secret() string {
	return c.secret
}

// Expiration returns the refresh token expiration
func (c *refreshTokenConfig) Expiration() time.Duration {
	return time.Minute * time.Duration(c.expiration)
}
