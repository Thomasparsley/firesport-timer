package main

import (
	"fmt"
	"strings"
	"time"

	Serial "github.com/tarm/serial"
	"thomasparsley.cz/firesport-timer/pkg/serial"
)

const (
	demoRawBuffer = "323a393133303a383a343432303a383a393133303a383a303a313a303a303a303a303a380d"
)

func StartConsole() {
	/*
		From string to buffer.
		This simulate buffer read from timer
	*/
	/* buffer := bytes.NewBufferString(demoRawBuffer)

	s := dual150.DecodeHexString(buffer.String())
	sSplit := strings.Split(s, ":")

	terminal.Clear()
	fmt.Printf("%s\n", sSplit)

	t, err := strconv.Atoi(sSplit[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Date(1, 1, 1, 0, 0, 0, t*1000000, time.Local))

	stringStatus, err := strconv.Atoi(sSplit[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(dual150.GetStatus(stringStatus))

	stringStatusa, err := strconv.Atoi(sSplit[3])
	if err != nil {
		panic(err)
	}
	fmt.Println(dual150.GetStatus(stringStatusa))

	stringStatusb, err := strconv.Atoi(sSplit[5])
	if err != nil {
		panic(err)
	}
	fmt.Println(dual150.GetStatus(stringStatusb)) */

	// 115200
	bauds := []int{ /* 4800, 9600, 19200, 38400, 57600, */ 115200 /* , 230400 */}

	for _, baud := range bauds {
		fmt.Println(baud)

		serialPortConfig := &Serial.Config{
			Name:        "COM4",
			Baud:        baud,
			ReadTimeout: time.Second * 5,
		}
		sa, err := Serial.OpenPort(serialPortConfig)
		if err != nil {
			panic(err)
		}
		fmt.Println("Port open")

		for i := 0; i < 9900; i++ {
			output, _ := ReadLine(sa, "#APP:cw:data?")

			fmt.Println(output)

			time.Sleep(time.Second / 30)
		}

		sa.Close()
		fmt.Println("Port close")
	}
}

/*
	2:0:1:0:1:0:1:0:1:0:0:0:0:1
	Neutrální stav

	2:0:4:6270:2:6270:2:0:1:0:0:0:0:2
	Levý a pravý terč běží

	2:0:8:32010:8:29470:8:0:1:0:0:0:0:8
	Finále
	Levý terč je výsledný čas

	2:0:4:5390:8:18070:2:0:1:0:0:0:0:2
	Běží pouze pravý terč, levý terč je sražen

	2:300000:1:0:1:0:1:0:1:0:0:0:0:1
	Nespuštění odpočet
*/

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
			panic(err)
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

/* package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	Serial "github.com/tarm/serial"
	"thomasparsley.cz/firesport-timer/internal/dual150"
	"thomasparsley.cz/firesport-timer/internal/terminal"
	"thomasparsley.cz/firesport-timer/pkg/serial"
)

const (
	demoRawBuffer = "323a393133303a383a343432303a383a393133303a383a303a313a303a303a303a303a380d"
)

func demo() {
		From string to buffer.
		This simulate buffer read from timer
	buffer := bytes.NewBufferString(demoRawBuffer)

	s := dual150.DecodeHexString(buffer.String())
	sSplit := strings.Split(s, ":")

	terminal.Clear()
	fmt.Printf("%s\n", sSplit)

	t, err := strconv.Atoi(sSplit[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Date(1, 1, 1, 0, 0, 0, t*1000000, time.Local))

	stringStatus, err := strconv.Atoi(sSplit[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(dual150.GetStatus(stringStatus))

	stringStatusa, err := strconv.Atoi(sSplit[3])
	if err != nil {
		panic(err)
	}
	fmt.Println(dual150.GetStatus(stringStatusa))

	stringStatusb, err := strconv.Atoi(sSplit[5])
	if err != nil {
		panic(err)
	}
	fmt.Println(dual150.GetStatus(stringStatusb))
}

func StartConsole() {
	// 115200
	bauds := []int{ 4800, 9600, 19200, 38400, 57600,115200, 230400}

	for _, baud := range bauds {
		fmt.Println(baud)

		serialPortConfig := &Serial.Config{
			Name:        "COM4",
			Baud:        baud,
			ReadTimeout: time.Second * 5,
		}
		ser, err := Serial.OpenPort(serialPortConfig)
		if err != nil {
			panic(err)
		}
		fmt.Println("Port open")

		for i := 0; i < 990; i++ {
			_, err = serial.WriteLine(ser, "#APP:cw:data?")
			if err != nil {
				panic(err)
			}
			fmt.Println("Write done")

			buf := make([]byte, 1024)
			n, _ := ser.Read(buf)
			fmt.Println("Read done")

			fmt.Printf("%q", buf[:n])
			fmt.Println()
		}

		ser.Close()
		fmt.Println("Port close")
	}
}
*/
