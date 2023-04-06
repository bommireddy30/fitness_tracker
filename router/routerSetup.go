package router

import (
	"Connection/handler"

	"github.com/gofiber/fiber/v2"
)

func RouterSetup(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create-user", handler.CreateUser)
}
