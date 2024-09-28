package auth

import (
	"github.com/gofiber/fiber/v2"
	"hackthehill/backend/database"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


func Register(c *fiber.Ctx) error {

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")


	db := database.GetDB()

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not hash password",
		})
	}

	user := database.User{
		"username": username,
		"email": email,
		"password": hashedPassword,
	}
	

	err = database.InsertUser(db, user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not create user",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User created",
	})

	


}