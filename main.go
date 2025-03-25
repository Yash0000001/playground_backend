package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/yash0000001/playgroundbackend/routes"
)

func main() {
	fmt.Println("Hello world")
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Middleware to add security headers
	app.Use(func(c *fiber.Ctx) error {
		// Add Cache-Control header
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")

		// Add X-Content-Type-Options header
		c.Set("X-Content-Type-Options", "nosniff")

		// Call the next middleware or handler
		return c.Next()
	})

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                  // Replace with your frontend domain
		AllowMethods: "POST, GET, OPTIONS", // Allowed HTTP methods
		AllowHeaders: "Content-Type",       // Allowed headers
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/v1/run", routes.Run_handler)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
