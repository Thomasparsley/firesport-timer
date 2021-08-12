package main

import (
	"github.com/gofiber/fiber/v2"

	"thomasparsley.cz/firesport-timer/routes"
)

func initRoutes(app *fiber.App, errorChan chan string, appVersion string) {
	routes.Homepage(app, appVersion)
	routes.Socket(app, errorChan)
}
