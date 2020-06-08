package custommiddleware

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nauhalf/mobilcantik/response"
)

func JWTUserCredential(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		user := c.Get("username").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		username := claims["username"]
		if username == nil {
			return c.JSON(http.StatusForbidden, response.ResponseError{
				Code:      http.StatusForbidden,
				Message:   http.StatusText(http.StatusForbidden),
				ErrorCode: nil,
			})
		}
		return next(c)
	}
}

func JWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("API_SECRET")),
		ContextKey: "username",
	})
}
