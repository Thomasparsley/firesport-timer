package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	Serial "github.com/tarm/serial"
	"github.com/zserge/lorca"

	"thomasparsley.cz/firesport-timer/internal/kocab"
	"thomasparsley.cz/firesport-timer/internal/serial"
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
		/* ticker := time.NewTicker(time.Second / 24).C */

		httpLinkChan := make(chan string)
		errorChan := make(chan string, 10)
		var httpLink string

		startReader := make(chan bool)
		closeReader := make(chan bool)
		portName := ""
		resetDual150Chan := make(chan bool)
		dual150Chan := make(chan kocab.Dual150, 2)
		dual150 := kocab.Dual150{}.New()

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
		for {
			if exitApp {
				break
			}

			// Start serial port
			select {
			case v, ok := <-startReader:
				if ok && v {
					go startSerialReader(portName, resetDual150Chan, errorChan, dual150Chan, closeReader)
				}
			case v, ok := <-dual150Chan:
				if ok {
					dual150 = v
				}
			case <-webui.Done():
				exitApp = true
			default:
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
				_, err := serial.WriteLine(sa, kocab.ResetDual150)
				if err != nil {
					errorChan <- err.Error()
				}

			}
		default:
			output, err := ReadLine(sa, kocab.ReadFromDual150)
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
