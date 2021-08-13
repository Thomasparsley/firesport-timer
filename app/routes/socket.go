package routes

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	websocket "github.com/gofiber/websocket/v2"
	"thomasparsley.cz/firesport-timer/internal/kocab"
)

func Socket(app *fiber.App, errorChan chan string, dual150 *kocab.Dual150) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		for {
			toSend := map[string]interface{}{
				"countdown": kocab.FormatTime(dual150.Countdown.Time),
				"lineOne":   kocab.FormatTime(dual150.LineOne.Time),
				"lineTwo":   kocab.FormatTime(dual150.LineTwo.Time),
				"lineThree": kocab.FormatTime(dual150.LineThree.Time),
				"lineFour":  kocab.FormatTime(dual150.LineFour.Time),
			}

			b, err := json.Marshal(toSend)
			if err != nil {
				errorChan <- fmt.Sprintf("Error: %s", err)
				break
			}

			err = c.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				errorChan <- fmt.Sprintf("Error: %s", err)
				break
			}

			time.Sleep(time.Second / 15)
		}
	}))
}
