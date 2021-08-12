package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Homepage(app *fiber.App, appVersion string) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		return c.SendString(fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="cs">

			<html>
			<head>
				<meta charset="UTF-8">
				<meta http-equiv="X-UA-Compatible" content="IE=edge">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">

				<link rel="preload" href="https://thomasparsley.cz/firesport-timer/index.css" as="style">
    			<link rel="stylesheet" href="https://thomasparsley.cz/firesport-timer/index.css" media="print" onload="this.media='all'">
    			<link rel="stylesheet" href="https://thomasparsley.cz/firesport-timer/index.css" type="text/css" as="style">

				<title>Firesport Timer</title>
			</head>

			<body data-app-version="%s">
				<div id="app"></div>

				<script src="https://thomasparsley.cz/firesport-timer/index.js"></script>
			</body>
			</html>
		`, appVersion))
	})
}
