package helpers

import "github.com/golang-jwt/jwt/v5"

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
	Userid   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type UserView struct {
	UserID   int         `json:"user_id,omitempty"`
	Username string      `json:"username,omitempty"`
	Email    string      `json:"email,omitempty"`
	Role     string      `json:"role,omitempty"`
	Comments interface{} `json:"comments,omitempty"`
	Posts    interface{} `json:"posts,omitempty"`
}
