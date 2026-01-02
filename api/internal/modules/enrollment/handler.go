package enrollment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) EnrollUser(c echo.Context) error {
	courseID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid course id"})
	}

	userID := c.QueryParam("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "user_id required"})
	}

	if err := h.repo.Enroll(c.Request().Context(), userID, courseID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to enroll"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"status": "enrolled"})
}
