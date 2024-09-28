package auth

import (
	"hackthehill/backend/ai"

	"github.com/gofiber/fiber/v2"

)

func TestEndpoint(c *fiber.Ctx) error {


	ai.TestAI()

	return c.SendString(c.Locals("user").(string))

	
}	