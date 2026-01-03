package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"github.com/NanoCode2022/online-courses-platform/api/internal/config"
	"github.com/NanoCode2022/online-courses-platform/api/internal/database"
	"github.com/NanoCode2022/online-courses-platform/api/internal/router"
)

func main() {
	// 1️⃣ Cargar configuración
	cfg := config.Load()

	// 2️⃣ Conectar MongoDB
	mongo := database.Connect(cfg.MongoURI, cfg.MongoDB)
	defer func() {
		if err := mongo.Client.Disconnect(nil); err != nil {
			log.Println("Error disconnecting Mongo:", err)
		}
	}()

	// 3️⃣ Crear Echo
	e := echo.New()

	// (opcional) middlewares globales
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// 4️⃣ Registrar rutas
	router.Register(
		e,
		mongo.DB,
		cfg.SupabaseJWKSURL,
	)

	// 5️⃣ Levantar servidor
	address := fmt.Sprintf(":%s", cfg.Port)
	log.Println("API running on", address)
	e.Logger.Fatal(e.Start(address))
}
