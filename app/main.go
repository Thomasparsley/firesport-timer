package main

import (
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/gofiber/fiber/v2"
	Serial "github.com/tarm/serial"
	"thomasparsley.cz/firesport-timer/internal/kocab"
	"thomasparsley.cz/firesport-timer/internal/serial"
	"thomasparsley.cz/firesport-timer/internal/terminal"
)

const (
	Dev = bool(false)

	pressToQuitApp   = string("ZMÁČKNĚTĚ `q` PRO UKONČENÍ APLIKACE")
	confirmToQuitApp = string("ZMÁČKNĚTĚ `y` PRO POTVRZENÍ UKONČENÍ NEBO `n` PRO ZRUŠENÍ UKONČENÍ")
)

func getActualTime() string {
	return time.Now().Format("15:04:05")
}

type portNameHttp struct {
	Port string `json:"port"`
}

func main() {
	if Dev {
		log.Println("[INFO] Development mode enabled")
	} else {
		if err := ui.Init(); err != nil {
			log.Fatal("failed to initialize termui:")
			panic(err)
		}
		defer ui.Close()

		uiEvents := ui.PollEvents()
		toExit := false

		tickerCount := 0
		tickerCount++
		ticker := time.NewTicker(time.Second / 24).C

		httpLinkChan := make(chan string)
		errorChan := make(chan string, 10)
		var httpLink string

		startReader := make(chan bool)
		closeReader := make(chan bool)
		portName := ""
		resetDual150Chan := make(chan bool)
		dual150Chan := make(chan kocab.Dual150, 2)
		dual150 := kocab.Dual150{}.New()

		// ===========
		// UI
		title := widgets.NewParagraph()
		title.Title = " Název aplikace: "
		title.Text = "Firesport Timer"
		title.SetRect(0, 0, 50, 3)

		author := widgets.NewParagraph()
		author.Title = " Autor: "
		author.Text = "https://thomasparsley.cz/"
		author.SetRect(49, 0, 100, 3)

		// Portname
		portNameUI := widgets.NewParagraph()
		portNameUI.Title = " Název portu: "
		portNameUI.Text = ""
		portNameUI.SetRect(39, 4, 80, 7)

		// Countdown
		countdown := widgets.NewParagraph()
		countdown.Title = " Odpočet: "
		countdown.Text = kocab.FormatTime(dual150.Countdown.Time)
		countdown.SetRect(0, 4, 40, 7)

		// Line one -- Left
		lineOne := widgets.NewParagraph()
		lineOne.Title = " Dráha první (VLEVO): "
		lineOne.Text = kocab.FormatTime(dual150.LineOne.Time)
		lineOne.SetRect(0, 7, 40, 10)

		// Line two -- Right
		lineTwo := widgets.NewParagraph()
		lineTwo.Title = " Dráha druhá (VPRAVO): "
		lineTwo.Text = kocab.FormatTime(dual150.LineTwo.Time)
		lineTwo.SetRect(39, 7, 80, 10)

		// Line three -- Left
		lineThree := widgets.NewParagraph()
		lineThree.Title = " Dráha třetí: "
		lineThree.Text = kocab.FormatTime(dual150.LineThree.Time)
		lineThree.SetRect(0, 10, 40, 13)

		// Line four -- Right
		lineFour := widgets.NewParagraph()
		lineFour.Title = " Dráha čtvrtá: "
		lineFour.Text = kocab.FormatTime(dual150.LineFour.Time)
		lineFour.SetRect(39, 10, 80, 13)

		// HTTP link
		httpLinkUI := widgets.NewParagraph()
		httpLinkUI.Text = httpLink
		httpLinkUI.Border = false
		httpLinkUI.SetRect(0, 13, 56, 16)

		// Quit
		quit := widgets.NewParagraph()
		quit.Text = pressToQuitApp
		quit.SetRect(0, 15, 45, 19)
		quit.TextStyle.Fg = ui.ColorWhite

		actualTime := widgets.NewParagraph()
		actualTime.Text = getActualTime()
		actualTime.SetRect(44, 15, 56, 19)

		// Error message
		errorMessage := widgets.NewParagraph()
		errorMessage.Text = ""
		errorMessage.Border = false
		errorMessage.TextStyle.Fg = ui.ColorRed
		errorMessage.SetRect(0, 18, 80, 22)
		// UI
		// ===========

		updateQuit := func(count int) {
			if count >= 0 && count <= 12 {
				quit.TextStyle.Fg = ui.ColorRed
			} else {
				quit.TextStyle.Fg = ui.ColorWhite
			}
		}

		draw := func() {
			actualTime.Text = getActualTime()

			httpLinkUI.Text = httpLink
			portNameUI.Text = portName

			// Error load
			select {
			case v, ok := <-errorChan:
				if ok {
					errorMessage.Text = v
				}
			default:
			}

			// Dual150 Load
			select {
			case v, ok := <-dual150Chan:
				if ok {
					dual150 = v

					countdown.Text = kocab.FormatTime(dual150.Countdown.Time)
					lineOne.Text = kocab.FormatTime(dual150.LineOne.Time)
					lineTwo.Text = kocab.FormatTime(dual150.LineTwo.Time)
					lineThree.Text = kocab.FormatTime(dual150.LineThree.Time)
					lineFour.Text = kocab.FormatTime(dual150.LineFour.Time)
				}
			default:
			}

			ui.Render(
				errorMessage,
				title,
				author,
				countdown,
				portNameUI,
				lineOne,
				lineTwo,
				lineThree,
				lineFour,
				httpLinkUI,
				quit,
				actualTime,
			)
		}

		// Start HTTP server
		go func() {
			app := http(errorChan, &dual150)

			app.Post("/api/start", func(c *fiber.Ctx) error {
				var port portNameHttp
				err := c.BodyParser(&port)
				if err != nil {
					errorChan <- err.Error()
					return c.SendStatus(500)
				} else if port.Port == "" {
					errorChan <- "Port is empty"
					return c.SendStatus(500)
				}

				portName = port.Port
				startReader <- true
				return c.SendStatus(200)
			})

			app.Post("/api/reset", func(c *fiber.Ctx) error {
				resetDual150Chan <- true
				return c.SendStatus(200)
			})

			app.Post("/api/close", func(c *fiber.Ctx) error {
				closeReader <- true
				portName = ""
				return c.SendStatus(200)
			})

			link := "127.0.0.1:3000"
			httpLinkChan <- "http://" + link + "/"
			app.Listen(link)
		}()
		httpLink = <-httpLinkChan

		terminal.Clear()
		draw()
		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					toExit = true
					quit.Text = confirmToQuitApp
				case "y":
					if toExit {
						return
					}
				case "n":
					if toExit {
						toExit = false
						quit.Text = pressToQuitApp
					}
				}

			case <-ticker:
				draw()
				updateQuit(tickerCount)
				tickerCount++

				if tickerCount%24 == 0 {
					tickerCount = 0
				}

				// Start serial port
				select {
				case v, ok := <-startReader:
					if ok && v {
						go startSerialReader(portName, resetDual150Chan, errorChan, dual150Chan, closeReader)
					}
				default:
				}
			}
		}
	}
}

func startSerialReader(portName string, resetDual150Chan chan bool, errorChan chan string, dualChan chan kocab.Dual150, closeReader chan bool) {
	serialPortConfig := &Serial.Config{
		Name:        portName, //"COM4",
		Baud:        115200,
		ReadTimeout: time.Second * 5,
	}

	sa, err := Serial.OpenPort(serialPortConfig)
	if err != nil {
		errorChan <- err.Error()
	}

	close := false
	defer func() {
		if r := recover(); r != nil {
			errorChan <- fmt.Sprintln("Recovered in startSerialReader", r)
		}
	}()

	for {
		select {
		case v, ok := <-closeReader:
			if ok && v {
				close = v
			}
		case v, ok := <-resetDual150Chan:
			if ok && v {
				_, err := serial.WriteLine(sa, "#RST")
				if err != nil {
					errorChan <- err.Error()
				}

			}
		default:
			output, err := ReadLine(sa, "#APP:cw:data?")
			if err != nil {
				errorChan <- err.Error()
				break
			}

			d, err := kocab.Dual150{}.ParseRawData(output)
			if err != nil {
				errorChan <- err.Error()
				continue
			}

			dualChan <- d

			time.Sleep(time.Second / 12)
		}

		if close {
			break
		}
	}

	sa.Close()
}
