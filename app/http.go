package main

import (
	"github.com/gofiber/fiber/v2"
	"thomasparsley.cz/firesport-timer/internal/kocab"
)

func http(errorChan chan string, dual150 *kocab.Dual150) *fiber.App {
	app := fiber.New()

	initMiddleware(app)
	initRoutes(app, errorChan, dual150)

	return app
}
