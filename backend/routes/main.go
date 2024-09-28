package routes

import (
	"github.com/gofiber/fiber/v2"
	"hackthehill/backend/database"
	"hackthehill/backend/auth"

	"hackthehill/backend/profiler"

)	


func Setup(app *fiber.App) {

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World from seconds file!")
	})


	app.Get("/test", func (c *fiber.Ctx) error {
		db := database.GetDB()
		user := database.User{
			"Username": "test",
			"Email": "test",
			"Password": "test",
		}
		database.InsertUser(db, user)
		return c.SendString("made")
	})

	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)
	// app.Get("/api/user", auth.UserData)

	api := app.Group("/api", auth.VerifyToken)
	api.Get("/check", auth.TestEndpoint)
	api.Post("/generateprofile", profiler.CreateProfile)


}