package routes

import (
	_ "embed"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//go:embed index.js
var js string

//go:embed index.css
var css string

func Homepage(app *fiber.App) {
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

			<body data-app-version="1.2.1" data-app-address="127.0.0.1" data-app-port="3000">
				<div id="app"></div>

				<style>%s</style>
				<script>%s</script>
			</body>
			</html>
		`, css, js))
	})
}
