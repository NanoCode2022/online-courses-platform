package lesson

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

type CreateLessonInput struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	VideoURL string `json:"video_url"`
	Order    int    `json:"order"`
}

func (h *Handler) Create(c echo.Context) error {
	courseIDParam := c.Param("id")

	courseID, err := primitive.ObjectIDFromHex(courseIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid course id",
		})
	}

	var input CreateLessonInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid payload",
		})
	}

	lesson := &Lesson{
		CourseID: courseID,
		Title:    input.Title,
		Content:  input.Content,
		VideoURL: input.VideoURL,
		Order:    input.Order,
	}

	if err := h.repo.Create(c.Request().Context(), lesson); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to create lesson",
		})
	}

	return c.JSON(http.StatusCreated, lesson)
}
