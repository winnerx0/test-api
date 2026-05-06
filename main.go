package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Deployed app on port " + port)
	})

	log.Printf("listening on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
