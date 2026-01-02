package course

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

func (h *Handler) GetAll(c echo.Context) error {
	courses, err := h.repo.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to fetch courses",
		})
	}

	return c.JSON(http.StatusOK, courses)
}

type CreateCourseInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (h *Handler) Create(c echo.Context) error {
	var input CreateCourseInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid payload",
		})
	}

	course := &Course{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Published:   false,
	}

	if err := h.repo.Create(c.Request().Context(), course); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to create course",
		})
	}

	return c.JSON(http.StatusCreated, course)
}

func (h *Handler) GetByID(c echo.Context) error {
	idParam := c.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid course id",
		})
	}

	course, err := h.repo.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "course not found",
		})
	}

	return c.JSON(http.StatusOK, course)
}
