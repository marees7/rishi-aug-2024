package middlewares

import (
	"blogs/api/validation"
	"blogs/common/dto"
	"blogs/pkg/loggers"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// verify if the user/admin has an valid token
func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//retrieve the stored token and data from the cookie
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			loggers.Warn.Println(err)
			return c.JSON(http.StatusUnauthorized, dto.ResponseJson{
				Message: "You need to login first to use blog post",
				Error:   err.Error(),
			})
		}

		//convert the retrieved token string into a jwt token
		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		
		if err != nil {
			loggers.Warn.Println(err)
			return c.JSON(http.StatusUnauthorized, dto.ResponseJson{
				Message: "invalid token",
				Error:   err.Error(),
			})
		}

		//retrieve the data stored inside token
		claims, err := validation.GetClaims(token)
		if err != nil {
			loggers.Warn.Println(err)
			return c.JSON(http.StatusUnauthorized, dto.ResponseJson{
				Message: "invalid token",
				Error:   err.Error(),
			})
		}

		//check if the token has expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.JSON(http.StatusUnauthorized, dto.ResponseJson{
				Message: "Session expired,please login again to continue",
			})
		}

		//set the values inside claims into the context
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])
		c.Set("email", claims["email"])

		return next(c)
	}
}
