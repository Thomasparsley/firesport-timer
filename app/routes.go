package main

import (
	"github.com/gofiber/fiber/v2"

	"thomasparsley.cz/firesport-timer/routes"
	"thomasparsley.cz/firesport-timer/timers/vendors/kocab"
)

func initRoutes(app *fiber.App, errorChan chan string, dual150 *kocab.Dual150) {
	routes.Homepage(app)
	routes.Socket(app, errorChan, dual150)

}
