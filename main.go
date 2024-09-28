package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"hackthehill/backend/routes"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
