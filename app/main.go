package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zserge/lorca"

	"thomasparsley.cz/firesport-timer/serialReader"
	"thomasparsley.cz/firesport-timer/timers/vendors/kocab"
)

const (
	Dev = bool(false)
)

type portNameHttp struct {
	Port string `json:"port"`
}

func main() {
	if Dev {
		/* log.Println("[INFO] Development mode enabled")

		sr := serialReader.New("COM4", 115200, time.Second)
		err := sr.Open()
		if err != nil {
			panic(err)
		}

		// for lopp in range 10
		i := 0
		for {
			sr.WriteLine(kocab.ReadFromDual150)
			l, err := sr.ReadLine()
			if err != nil {
				panic(err)
			}
			fmt.Println(l)

			time.Sleep(time.Second)

			if i > 10 {
				break
			}

			i++
		}

		sr.Close() */
	} else {
		var sr serialReader.Serial
		errorChan := make(chan string, 10)
		if r := recover(); r != nil {
			errorChan <- fmt.Sprintln("Recovered in main func", r)
		}

		defer func() {
			if sr.Config && sr.PortOpen {
				sr.Close()
			}
		}()

		httpLinkChan := make(chan string)
		var httpLink string

		startReader := make(chan bool)
		closeReader := make(chan bool)
		portName := ""
		var resetDual150 bool
		dual150 := kocab.NewDual150()

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
				resetDual150 = true
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

		webui, _ := lorca.New(httpLink, "", 1280, 720)
		defer webui.Close()

		var exitApp bool
		for {
			if exitApp {
				break
			}

			if sr.Config && sr.PortOpen {
				if resetDual150 {
					_, err := sr.WriteLine(kocab.ResetDual150)
					if err != nil {
						errorChan <- err.Error()
					}

					fmt.Println(resetDual150)
					resetDual150 = false
					fmt.Println(resetDual150)
				}

				_, err := sr.WriteLine(kocab.ReadFromDual150)
				if err != nil {
					errorChan <- err.Error()
					continue
				}

				output, err := sr.ReadLine()
				if err != nil {
					errorChan <- err.Error()
					continue
				}

				d, err := kocab.ParseDual150(output)
				if err != nil {
					errorChan <- err.Error()
					continue
				}

				dual150 = d
			}

			// Start serial port
			select {
			case v, ok := <-startReader:
				if ok && v {
					sr = serialReader.New(portName, 115200, time.Second)
					err := sr.Open()
					if err != nil {
						errorChan <- err.Error()
						continue
					}
				}
			case v, ok := <-closeReader:
				if ok && v {
					err := sr.Close()
					if err != nil {
						errorChan <- err.Error()
						continue
					}
				}
			case <-webui.Done():
				exitApp = true
			default:
			}

			time.Sleep(time.Second / 24)
		}
	}
}
