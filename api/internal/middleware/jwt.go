package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "missing token",
				})
			}

			tokenString := auth[len("Bearer "):]

			token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "invalid token",
				})
			}

			claims := token.Claims.(jwt.MapClaims)
			role := "user"

			if meta, ok := claims["user_metadata"].(map[string]interface{}); ok {
				if r, ok := meta["role"].(string); ok {
					role = r
				}
			}

			c.Set("user_role", role)

			// datos útiles
			c.Set("user_email", claims["email"])
			c.Set("user_id", claims["sub"])

			return next(c)
		}
	}
}
