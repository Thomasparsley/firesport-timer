package main

import (
	"bytes"
	"thomasparsley.cz/firesport-timer/bufferTransfare"
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

	s := bufferTransfare.DecodeHexString(buffer.String())
	println(s)
}
