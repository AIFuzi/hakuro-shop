package main

import (
	"github.com/AIFuzi/hakuro-shop/internal/database"
	"github.com/AIFuzi/hakuro-shop/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	database.New()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("ORIGINS"),
		AllowCredentials: true,
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
