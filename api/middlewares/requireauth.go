package middlewares

import (
	"blogs/common/helpers"
	"blogs/pkg/loggers"
	"blogs/api/validation"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			loggers.WarningLog.Println(err)
			return c.JSON(http.StatusRequestTimeout, helpers.ResponseJson{
				Message: "You need to login first to use blog post",
				Error:   err.Error(),
			})
		}

		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			loggers.WarningLog.Println(err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseJson{
				Message: "invalid token",
				Error:   err.Error(),
			})
		}

		claims, err := validation.GetClaims(token)
		if err != nil {
			loggers.WarningLog.Println(err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseJson{
				Message: "invalid token",
				Error:   err.Error(),
			})
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.JSON(http.StatusGatewayTimeout, helpers.ResponseJson{
				Message: "Session expired,please login again to continue",
			})
		} else if claims["user_id"] == 0 {
			loggers.WarningLog.Println(err)
			return c.JSON(http.StatusNotFound, helpers.ResponseJson{
				Message: "user not found",
			})
		}

		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		return next(c)
	}
}
