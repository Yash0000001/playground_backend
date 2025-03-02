package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/v1/run", routes.Run_handler)
	log.Fatal(app.Listen(":" + port))
}
