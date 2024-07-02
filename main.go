package main

import (
	"log"

	"taskday7.4/configs"
	"taskday7.4/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Memuat pengaturan dari file .env
	settings := configs.ImportSetting()

	// Inisialisasi database
	db, err := configs.ConnectDB(settings)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	configs.DB = db

	// Inisialisasi Echo
	e := echo.New()

	// Inisialisasi routes
	routes.RegisterRoutes(e)

	// Jalankan server
	e.Logger.Fatal(e.Start(":8080"))
}
