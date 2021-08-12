package main

import (
	"github.com/gofiber/fiber/v2"

	"thomasparsley.cz/firesport-timer/internal/kocab"
	"thomasparsley.cz/firesport-timer/routes"
)

func initRoutes(app *fiber.App, errorChan chan string, dualChan chan kocab.Dual150) {
	routes.Homepage(app)
	routes.Socket(app, errorChan, dualChan)

}
