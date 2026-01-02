package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole, ok := c.Get("user_role").(string)
			if !ok || userRole != role {
				return c.JSON(http.StatusForbidden, echo.Map{
					"error": "forbidden",
				})
			}
			return next(c)
		}
	}
}
