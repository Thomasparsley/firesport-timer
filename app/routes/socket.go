package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	websocket "github.com/gofiber/websocket/v2"
	"thomasparsley.cz/firesport-timer/internal/dual150"
)

const (
	demoRawBuffer = "323a393133303a383a343432303a383a393133303a383a303a313a303a303a303a303a380d"
)

type timer struct {
	Left        time.Time      `json:"left"`
	LeftStatus  dual150.Status `json:"leftStatus"`
	Right       time.Time      `json:"right"`
	RightStatus dual150.Status `json:"rightStatus"`
}

func Socket(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	/* app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host"))

		for {
			buffer := bytes.NewBufferString(demoRawBuffer)

			s := dual150.DecodeHexString(buffer.String())

			sSplit := strings.Split(s, ":")

			t, err := strconv.Atoi(sSplit[0])
			if err != nil {
				panic(err)
			}

			tt := time.Date(1, 1, 1, 0, 0, 0, t*1000000, time.Local)

			stringStatusa, err := strconv.Atoi(sSplit[3])
			if err != nil {
				panic(err)
			}

			stringStatusb, err := strconv.Atoi(sSplit[5])
			if err != nil {
				panic(err)
			}

			tim := timer{
				Left:        "tt",
				LeftStatus:  dual150.GetStatus(stringStatusa),
				Right:       "tt",
				RightStatus: dual150.GetStatus(stringStatusb),
			}

			b, err := json.Marshal(tim)
			if err != nil {
				fmt.Printf("Error: %s", err)
				break
			}

			fmt.Println(string(b))

			err = c.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Println("read:", err)
				break
			}

			time.Sleep(time.Second / 10)
		}
	})) */
}
