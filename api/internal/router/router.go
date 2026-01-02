package router

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"

	courseModule "github.com/NanoCode2022/online-courses-platform/api/internal/modules/course"
	enrollmentModule "github.com/NanoCode2022/online-courses-platform/api/internal/modules/enrollment"
	lessonModule "github.com/NanoCode2022/online-courses-platform/api/internal/modules/lesson"

	authMiddleware "github.com/NanoCode2022/online-courses-platform/api/internal/middleware"
)

func Register(
	e *echo.Echo,
	db *mongo.Database,
	jwtSecret string,
) {
	// =====================
	// Infraestructura
	// =====================

	// Healthcheck
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"status": "ok",
		})
	})

	// =====================
	// Repositorios
	// =====================

	courseRepo := courseModule.NewMongoRepository(db)
	lessonRepo := lessonModule.NewMongoRepository(db)
	enrollmentRepo := enrollmentModule.NewMongoRepository(db)

	// =====================
	// Handlers
	// =====================

	courseHandler := courseModule.NewHandler(courseRepo)
	lessonHandler := lessonModule.NewHandler(lessonRepo)
	enrollmentHandler := enrollmentModule.NewHandler(enrollmentRepo)

	// =====================
	// Rutas públicas
	// =====================

	e.GET("/courses", courseHandler.GetAll)
	e.GET("/courses/:id", courseHandler.GetByID)

	// =====================
	// Rutas autenticadas
	// =====================

	auth := e.Group("/api")
	auth.Use(authMiddleware.JWT(jwtSecret))

	// endpoint de prueba
	auth.GET("/me", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"email": c.Get("user_email"),
			"role":  c.Get("user_role"),
		})
	})

	// =====================
	// Rutas ADMIN
	// =====================

	admin := auth.Group("/admin")
	admin.Use(authMiddleware.RequireRole("admin"))

	// Cursos
	admin.POST("/courses", courseHandler.Create)

	// Lecciones
	admin.POST("/courses/:id/lessons", lessonHandler.Create)

	// Inscripciones (simulación de compra)
	admin.POST("/courses/:id/enroll", enrollmentHandler.EnrollUser)

	// =====================
	// Rutas de ALUMNO (contenido protegido)
	// =====================

	student := auth.Group("/courses")
	student.GET("/:id/lessons", func(c echo.Context) error {
		// Este endpoint normalmente:
		// 1. verifica inscripción
		// 2. devuelve lecciones
		// (la lógica vive en el handler real)
		return c.JSON(200, echo.Map{
			"message": "lessons endpoint (protected)",
		})
	})
}
