package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3007"
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
		})
	})

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error in Listen: %s", err)
	}
}
