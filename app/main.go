package main

import (
	"fmt"
	"log"
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
		log.Println("[INFO] Development mode enabled")
	} else {
		httpLinkChan := make(chan string)
		errorChan := make(chan string, 10)
		var httpLink string

		startReader := make(chan bool)
		closeReader := make(chan bool)
		portName := ""
		resetDual150Chan := make(chan bool)
		dual150 := kocab.NewDual150()

		if r := recover(); r != nil {
			errorChan <- fmt.Sprintln("Recovered in main func", r)
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

		webui, _ := lorca.New(httpLink, "", 1280, 720)
		defer webui.Close()

		var exitApp bool
		var sr serialReader.Serial
		for {
			if exitApp {
				break
			}

			if sr.Config && sr.PortOpen {
				select {
				case v, ok := <-resetDual150Chan:
					if ok && v {
						_, err := sr.WriteLine(kocab.ResetDual150)
						if err != nil {
							errorChan <- err.Error()
							continue
						}
					}
				default:
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
					}
				}
			case v, ok := <-closeReader:
				if ok && v {
					err := sr.Close()
					if err != nil {
						errorChan <- err.Error()
					}
				}
			case <-webui.Done():
				exitApp = true
			default:
			}

		}
	}
}
