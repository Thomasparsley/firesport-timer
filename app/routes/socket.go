package routes

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	websocket "github.com/gofiber/websocket/v2"
	"thomasparsley.cz/firesport-timer/internal/kocab"
)

func Socket(app *fiber.App, errorChan chan string) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		for {

			dual, err := kocab.Dual150{}.ParseRawData("2:300000:1:0:1:0:1:0:1:0:0:0:0:1")
			if err != nil {
				errorChan <- fmt.Sprintf("Error: %s", err)
				break
			}

			toSend := map[string]interface{}{
				"countdown": kocab.FormatTime(dual.Countdown.Time),
				"lineOne":   kocab.FormatTime(dual.LineOne.Time),
				"lineTwo":   kocab.FormatTime(dual.LineTwo.Time),
				"lineThree": kocab.FormatTime(dual.LineThree.Time),
				"lineFour":  kocab.FormatTime(dual.LineFour.Time),
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
