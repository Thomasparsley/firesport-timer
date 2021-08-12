package main

import (
	"github.com/gofiber/fiber/v2"
)

func http(errorChan chan string, appVersion string) *fiber.App {
	app := fiber.New()

	initMiddleware(app)
	initRoutes(app, errorChan, appVersion)

	return app
}
