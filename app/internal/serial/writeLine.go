package serial

import Serial "github.com/tarm/serial"

func WriteLine(port *Serial.Port, data string) (int, error) {
	return port.Write([]byte(data + "\n"))
}
