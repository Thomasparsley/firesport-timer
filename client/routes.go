package main

import (
	"github.com/gofiber/fiber/v2"

	"thomasparsley.cz/firesport-timer/routes"
)

func initRoutes(app *fiber.App) {
	routes.Homepage(app)
}
