package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ResponseJson struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaims struct {
	Userid   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}

type UserView struct {
	UserID   uuid.UUID   `json:"user_id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Role     string      `json:"role"`
	Comments interface{} `json:"comments"`
	Posts    interface{} `json:"posts"`
}