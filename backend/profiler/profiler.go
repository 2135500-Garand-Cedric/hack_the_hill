package profiler

import (
	"fmt"
	"hackthehill/backend/database"

	"github.com/gofiber/fiber/v2"
)


func CreateProfile(c *fiber.Ctx) error {

	goals := c.FormValue("goals")
	hobbies := c.FormValue("hobbies")
	interests := c.FormValue("interests")
	Occupation := c.FormValue("occupation")
	dob := c.FormValue("dob")
	gender := c.FormValue("gender")
	city := c.FormValue("city")

	username := c.Locals("user").(string)

	fmt.Println(username)

	db := database.GetDB()
	user, err := database.FindUserByUsername(db, username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not find user",
		})
	}

	profile := database.Profile{
		"username": user["username"],
		"goals": goals,
		"hobbies": hobbies,
		"interests": interests,
		"occupation": Occupation,
		"dob": dob,
		"gender": gender,
		"city": city,
	}

	db = database.GetProfileDB()
	err = database.InsertProfile(db, profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Profile created",
	})
}

func GetProfile(c *fiber.Ctx) error {

	db := database.GetProfileDB()
	profile, err := database.GetProfile(db, c.Locals("user").(string))
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(profile)
}