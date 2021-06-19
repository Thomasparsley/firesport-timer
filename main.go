package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tarm/serial"

	"thomasparsley.cz/firesport-timer/dual150"
	"thomasparsley.cz/firesport-timer/terminal"
)

const (
	demoRawBuffer = "323a393133303a383a343432303a383a393133303a383a303a313a303a303a303a303a380d"
)

func main() {
	/*
		From string to buffer.
		This simulate buffer read from timer
	*/
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

	bauds := []int{ /* 4800, 9600, 19200, 38400, 57600,  */ 115200 /* , 230400 */}

	for _, baud := range bauds {
		fmt.Println(baud)

		serialPortConfig := &serial.Config{
			Name:        "COM4",
			Baud:        baud,
			ReadTimeout: time.Second * 5,
			/* Size:        8, */
		}
		sa, err := serial.OpenPort(serialPortConfig)
		if err != nil {
			panic(err)
		}
		fmt.Println("Port open")

		for i := 0; i < 9900; i++ {
			_, err = sa.Write([]byte("#APP:cw:data?\n"))
			if err != nil {
				panic(err)
			}
			fmt.Println("Write done")

			buf := make([]byte, 1024)
			n, _ := sa.Read(buf)
			fmt.Println("Read done")

			fmt.Printf("%q", buf[:n])
			fmt.Println()
		}

		sa.Close()
		fmt.Println("Port close")
	}

}
