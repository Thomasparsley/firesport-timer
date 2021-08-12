package main

import (
	"github.com/gofiber/fiber/v2"
	"thomasparsley.cz/firesport-timer/internal/kocab"
)

func http(errorChan chan string, dualChan chan kocab.Dual150) *fiber.App {
	app := fiber.New()

	initMiddleware(app)
	initRoutes(app, errorChan, dualChan)

	return app
}
