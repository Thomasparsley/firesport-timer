package main

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	Serial "github.com/tarm/serial"
	"thomasparsley.cz/firesport-timer/internal/kocab"
	"thomasparsley.cz/firesport-timer/internal/terminal"
)

const (
	appVersion = string("1.0.1")

	Dev = bool(false)

	pressToQuitApp   = string("ZMÁČKNĚTĚ `q` PRO UKONČENÍ APLIKACE")
	confirmToQuitApp = string("ZMÁČKNĚTĚ `y` PRO POTVRZENÍ UKONČENÍ NEBO `n` PRO ZRUŠENÍ UKONČENÍ")
)

func getActualTime() string {
	return time.Now().Format("15:04:05")
}

func main() {
	if Dev {
		log.Println("[INFO] Development mode enabled")
		StartConsole()
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

		title := widgets.NewParagraph()
		title.Title = " Název aplikace: "
		title.Text = "Firesport Timer"
		title.SetRect(0, 0, 50, 3)

		author := widgets.NewParagraph()
		author.Title = " Autor: "
		author.Text = "https://thomasparsley.cz/"
		author.SetRect(49, 0, 100, 3)

		closeReader := make(chan bool)
		dualChan := make(chan kocab.Dual150, 2)
		dual := kocab.Dual150{}.New()
		/* dual, err := kocab.Dual150{}.ParseRawData("2:300000:1:0:1:0:1:0:1:0:0:0:0:1")
		if err != nil {
			panic(err)
		} */

		// Countdown
		countdown := widgets.NewParagraph()
		countdown.Title = " Odpočet: "
		countdown.Text = kocab.FormatTime(dual.Countdown.Time)
		countdown.SetRect(0, 4, 40, 7)

		// Line one -- Left
		lineOne := widgets.NewParagraph()
		lineOne.Title = " Dráha první (VLEVO): "
		lineOne.Text = kocab.FormatTime(dual.LineOne.Time)
		lineOne.SetRect(0, 7, 40, 10)

		// Line two -- Right
		lineTwo := widgets.NewParagraph()
		lineTwo.Title = " Dráha druhá (VPRAVO): "
		lineTwo.Text = kocab.FormatTime(dual.LineTwo.Time)
		lineTwo.SetRect(39, 7, 80, 10)

		// Line three -- Left
		lineThree := widgets.NewParagraph()
		lineThree.Title = " Dráha třetí: "
		lineThree.Text = kocab.FormatTime(dual.LineThree.Time)
		lineThree.SetRect(0, 10, 40, 13)

		// Line four -- Right
		lineFour := widgets.NewParagraph()
		lineFour.Title = " Dráha čtvrtá: "
		lineFour.Text = kocab.FormatTime(dual.LineFour.Time)
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
			case v, ok := <-dualChan:
				if ok {
					dual = v

					countdown.Text = kocab.FormatTime(dual.Countdown.Time)
					lineOne.Text = kocab.FormatTime(dual.LineOne.Time)
					lineTwo.Text = kocab.FormatTime(dual.LineTwo.Time)
					lineThree.Text = kocab.FormatTime(dual.LineThree.Time)
					lineFour.Text = kocab.FormatTime(dual.LineFour.Time)
				}
			default:
			}

			ui.Render(errorMessage, title, author, countdown, lineOne, lineTwo, lineThree, lineFour, httpLinkUI, quit, actualTime)
		}

		// Start HTTP server
		go func() {
			app := http(errorChan, appVersion, dualChan)

			link := "127.0.0.1:3000"
			httpLinkChan <- "http://" + link + "/"
			app.Listen(link)
		}()
		httpLink = <-httpLinkChan

		// Start serial reader
		go func() {
			serialPortConfig := &Serial.Config{
				Name:        "COM4",
				Baud:        115200,
				ReadTimeout: time.Second * 5,
			}

			sa, err := Serial.OpenPort(serialPortConfig)
			if err != nil {
				panic(err)
			}

			close := false
			for {
				output, _ := ReadLine(sa, "#APP:cw:data?")

				d, err := kocab.Dual150{}.ParseRawData(output)
				if err != nil {
					errorChan <- err.Error()
					continue
				}

				dualChan <- d

				time.Sleep(time.Second / 12)

				select {
				case v, ok := <-closeReader:
					if ok {
						close = v
					}
				default:
				}

				if close {
					break
				}
			}

			sa.Close()
		}()

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
			}
		}
	}
}
