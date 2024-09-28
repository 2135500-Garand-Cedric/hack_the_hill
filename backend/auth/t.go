
package auth

import (
	"github.com/gofiber/fiber/v2"
)

func TestEndpoint(c *fiber.Ctx) error {

	


	return c.SendString(c.Locals("user").(string))

	
}	