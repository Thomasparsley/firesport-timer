package main

import (
	"log"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	Dev = bool(true)
)

func main() {
	if Dev {
		StartConsole()
		/* app := http()

		app.Listen("127.0.0.1:3000") */
	} else {
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		title := widgets.NewParagraph()
		title.Title = " App name: "
		title.Text = " Firesport Timer"
		title.SetRect(0, 0, 50, 3)

		author := widgets.NewParagraph()
		author.Title = " Author: "
		author.Text = " https://thomasparsley.cz/"
		author.SetRect(49, 0, 100, 3)

		p := widgets.NewParagraph()
		p.Text = time.Now().String()
		p.SetRect(0, 2, 100, 5)

		quit := widgets.NewParagraph()
		quit.Text = " PRESS `q` TO QUIT APP "
		quit.SetRect(0, 15, 35, 18)
		quit.TextStyle.Fg = ui.ColorWhite

		fps := widgets.NewParagraph()
		fps.Text = "0"
		fps.SetRect(34, 15, 40, 18)

		updateQuit := func(count int) {
			if count >= 0 && count <= 12 {
				quit.TextStyle.Fg = ui.ColorRed
			} else {
				quit.TextStyle.Fg = ui.ColorWhite
			}
		}

		draw := func(count int) {
			p.Text = time.Now().String()
			fps.Text = strconv.Itoa(count)

			ui.Render(title, author, p, quit, fps)
		}

		tickerCount := 0
		draw(tickerCount)
		tickerCount++

		uiEvents := ui.PollEvents()
		ticker := time.NewTicker(time.Second / 24).C
		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					return
				}
			case <-ticker:
				draw(tickerCount)
				updateQuit(tickerCount)
				tickerCount++
			}

			if tickerCount%24 == 0 {
				tickerCount = 0
			}
		}
	}
}
