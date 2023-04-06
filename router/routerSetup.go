package router

import (
	"Connection/handler"

	"github.com/gofiber/fiber/v2"
)

func RouterSetup(app *fiber.App) {
	api := app.Group("/api")
	user := api.Group("/user")
	user.Post("/create-user", handler.CreateUser)
}
