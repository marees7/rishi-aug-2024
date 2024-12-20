package validation

import (
	"blogs/common/dto"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// generate a new token for the user
func GenerateToken(user *models.User) (string, error) {
	//set claims with needed data and expire time if needed
	claims := &dto.JWTClaims{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 1800)),
		},
	}

	//creates a jwt token with the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//creates a token string
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		loggers.Warn.Println(err)
		return "", err
	}

	return tokenStr, nil
}

// retrieve the data inside the claims
func GetClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	
	return claims, nil
}
