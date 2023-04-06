package main

import (
	"Connection/database"
	"Connection/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	router.RouterSetup(app)

	app.Listen(":4200")
}
