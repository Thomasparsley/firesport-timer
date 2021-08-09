package routes

import "github.com/gofiber/fiber/v2"

func Homepage(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		return c.SendString(`
			<!DOCTYPE html>
			<html lang="cs">

			<html>
			<head>
				<meta charset="UTF-8">
				<meta http-equiv="X-UA-Compatible" content="IE=edge">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">

				<title>Firesport Timer</title>
			</head>

			<body>
				<div id="app"></div>
			</body>
			</html>
		`)
	})
}
