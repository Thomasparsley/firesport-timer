package main

import (
	"strings"

	Serial "github.com/tarm/serial"
	"thomasparsley.cz/firesport-timer/internal/serial"
)

/*
	Read line from serial port
*/
func ReadLine(sa *Serial.Port, write string) (string, error) {
	var (
		buf       []byte
		bufString string
	)

	for {
		_, err := serial.WriteLine(sa, write)
		if err != nil {
			return "", err
		}

		buf = make([]byte, 1024)
		n, err := sa.Read(buf)
		if err != nil {
			return "", err
		}

		bufString += string(buf[:n])

		if strings.Contains(bufString, "\r") {

			if len(bufString) < 15 {
				bufString = ""
			} else {
				if bufString[1] != ':' {
					bufString = ""
				} else {
					return bufString, nil
				}
			}
		}
	}
}
