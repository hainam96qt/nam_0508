package model

import (
	"time"
)

type Identity struct {
	Username string `json:"username"`
	USerID   int    `json:"user_id"`
}

type Success struct {
	Message string `json:"message"`
}

type TokenPair struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`

	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

type CreateRegistrationRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type CreateRegistrationResponse struct {
	TokenPair
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	TokenPair
}
