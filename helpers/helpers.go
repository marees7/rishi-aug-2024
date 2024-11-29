package helpers

import "github.com/golang-jwt/jwt/v5"

type ResponseJson struct {
	Message string      `json:"message"`
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
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Role     string      `json:"role"`
	Comments interface{} `json:"comments"`
	Posts    interface{} `json:"posts"`
}
