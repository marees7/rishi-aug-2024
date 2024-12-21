package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JSON response for every success and error response
type ResponseJson struct {
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        string      `json:"error,omitempty"`
	Limit        int         `json:"limit,omitempty"`
	Offset       int         `json:"offset,omitempty"`
	TotalRecords int64       `json:"total_records,omitempty"`
}

// Error response with http status code and error message
type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

// for login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// assign JWT claims along with registered claims
type JWTClaims struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}
