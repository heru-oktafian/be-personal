package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/heru-oktafian/be-personal/config"
	"github.com/heru-oktafian/be-personal/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Tidak menemukan file .env, menggunakan environment variables sistem.")
	}

	// Inisialisasi Database
	db := config.ConnectDB()
	defer db.Close()

	// Inisialisasi Fiber
	app := fiber.New(fiber.Config{
		AppName: "HeruOktafian CMS API",
	})

	// Middleware Global
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:5173", // Rails publik & React admin
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// 1. Setup Static Files (Folder tempat gambar disimpan)
	// Fiber akan melayani file di dalam "./public/uploads" ketika ada request ke "/uploads"
	app.Static("/uploads", "./public/uploads")

	// Setup API Routing
	// Hapus placeholder routing manual sebelumnya dan ganti dengan ini:
	routes.SetupRoutes(app, db)

	// Rute Publik
	public := app.Group("/public")
	public.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "API Backend Berjalan Baik"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di port %s", port)
	log.Fatal(app.Listen(":" + port))
}
