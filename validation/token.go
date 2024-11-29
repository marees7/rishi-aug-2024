package validation

import (
	"blogs/helpers"
	"blogs/loggers"
	"blogs/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.Users) (string, error) {
	claims := &helpers.JWTClaims{
		Userid:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 1200)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstr, err := token.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		loggers.WarningLog.Println(err)
		return "", err
	}

	return tokenstr, nil
}

func GetClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
