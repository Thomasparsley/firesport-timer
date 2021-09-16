package serialReader

import (
	"errors"
	"strings"
	"time"

	"github.com/tarm/serial"
)

type Serial struct {
	Name        string
	Baud        int
	ReadTimeout time.Duration

	Config   bool
	port     *serial.Port
	PortOpen bool
}

func New(name string, baud int, readTimeout time.Duration) Serial {
	return Serial{
		Name:        name,
		Baud:        baud,
		ReadTimeout: readTimeout,

		Config: true,
	}
}

func (s *Serial) Open() error {
	if !s.Config {
		return errors.New("serial is not configured")
	} else if s.port != nil {
		return errors.New("serial is already open")
	}

	var err error

	s.port, err = serial.OpenPort(&serial.Config{
		Name:        s.Name,
		Baud:        s.Baud,
		ReadTimeout: s.ReadTimeout,
	})
	if err != nil {
		return err
	}

	s.PortOpen = true
	return nil
}

func (s *Serial) Close() error {
	if !s.Config {
		return errors.New("serial is not configured")
	} else if s.port == nil {
		return errors.New("serial is not open")
	}

	s.PortOpen = false
	return s.port.Close()
}

func (s *Serial) Write(data string) (int, error) {
	if !s.Config {
		return 0, errors.New("serial is not configured")
	} else if s.port == nil {
		return 0, errors.New("serial is not open")
	}

	return s.port.Write([]byte(data))
}

func (s *Serial) WriteLine(data string) (int, error) {
	return s.Write(data + "\n")
}

func (s *Serial) Read() (string, error) {
	if !s.Config {
		return "", errors.New("serial is not configured")
	} else if s.port == nil {
		return "", errors.New("serial is not open")
	}

	buf := make([]byte, 512)
	n, err := s.port.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func (s *Serial) ReadLine() (string, error) {
	var stringBuff string
	var counter int

	for {
		readedLine, err := s.Read()
		if err != nil {
			return "", err
		}

		if strings.Contains(readedLine, "\r") {
			counter++
			stringBuff = strings.Split(stringBuff, "\r")[counter]

			if counter == 1 {
				break
			}
		} else {
			stringBuff += readedLine
		}
	}
	return stringBuff, nil
}
