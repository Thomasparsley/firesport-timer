package serial

import (
	Serial "github.com/tarm/serial"
)

func ReadLine(port *Serial.Port) []byte {
	var result []byte
	buffer := make([]byte, 1024)

	port.Read(buffer)

	return result
}
