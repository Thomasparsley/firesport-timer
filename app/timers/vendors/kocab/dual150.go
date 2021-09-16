package kocab

/*
	2:0:1:0:1:0:1:0:1:0:0:0:0:1
	Neutral state

	2:0:4:6270:2:6270:2:0:1:0:0:0:0:2
	Left and right target running

	2:0:8:32010:8:29470:8:0:1:0:0:0:0:8
	Final
	The left target is the resulting time

	2:0:4:5390:8:18070:2:0:1:0:0:0:0:2
	Only the right target runs, the left target is knocked down

	2:300000:1:0:1:0:1:0:1:0:0:0:0:1
	Countdown
*/

import (
	"errors"
	"log"
	"strings"

	"thomasparsley.cz/firesport-timer/timers/line"
)

const (
	ResetDual150    = string("#RST")
	ReadFromDual150 = string("#APP:cw:data?")
)

type Dual150 struct {
	Countdown line.Line `json:"countdown"` //
	LineOne   line.Line `json:"lineOne"`   // Left
	LineTwo   line.Line `json:"lineTwo"`   // Right
	LineThree line.Line `json:"lineThree"` //
	LineFour  line.Line `json:"lineFour"`  //
}

func (d Dual150) OneOfLinesIsZero() bool {
	return d.LineOne.IsZero() || d.LineTwo.IsZero() || d.LineThree.IsZero() || d.LineFour.IsZero()
}

func NewDual150() Dual150 {
	return Dual150{
		Countdown: line.New(),
		LineOne:   line.New(),
		LineTwo:   line.New(),
		LineThree: line.New(),
		LineFour:  line.New(),
	}
}

func ParseDual150(rawData string) (Dual150, error) {
	log.Println(rawData)
	if len(rawData) < 27 {
		return Dual150{}, errors.New("invalid rawData length")
	}

	// Check if rawData contain a 13x ':', if not, it's not a valid Dual150
	if strings.Count(rawData, ":") != 13 {
		log.Println("invalid rawData")
		return Dual150{}, errors.New("invalid rawData")
	}

	rawDataSplit := strings.Split(rawData, ":")

	var countdown line.Line

	lineOne, err := line.Parse(rawDataSplit[3], rawDataSplit[4])
	if err != nil {
		return Dual150{}, err
	}

	lineTwo, err := line.Parse(rawDataSplit[5], rawDataSplit[6])
	if err != nil {
		return Dual150{}, err
	}

	lineThree, err := line.Parse(rawDataSplit[7], rawDataSplit[8])
	if err != nil {
		return Dual150{}, err
	}

	lineFour, err := line.Parse(rawDataSplit[9], rawDataSplit[10])
	if err != nil {
		return Dual150{}, err
	}

	if !lineOne.IsZero() ||
		!lineTwo.IsZero() ||
		!lineThree.IsZero() ||
		!lineFour.IsZero() {
		countdown = line.New()
	} else {
		countdown, err = line.ParseCountdown(rawDataSplit[1])
		if err != nil {
			return Dual150{}, err
		}

		lineOne.SetDefault()
		lineTwo.SetDefault()
		lineThree.SetDefault()
		lineFour.SetDefault()
	}

	return Dual150{
		Countdown: countdown,
		LineOne:   lineOne,
		LineTwo:   lineTwo,
		LineThree: lineThree,
		LineFour:  lineFour,
	}, nil
}
