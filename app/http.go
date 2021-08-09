package main

import "github.com/gofiber/fiber/v2"

func http() *fiber.App {
	app := fiber.New()

	initMiddleware(app)
	initRoutes(app)

	return app
}
