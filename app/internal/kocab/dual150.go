package kocab

import (
	"errors"
	"strings"
)

const (
	ResetDual150    = string("#RST")
	ReadFromDual150 = string("#APP:cw:data?")
)

type Dual150 struct {
	Countdown Line `json:"countdown"` //
	LineOne   Line `json:"lineOne"`   // Left
	LineTwo   Line `json:"lineTwo"`   // Right
	LineThree Line `json:"lineThree"` //
	LineFour  Line `json:"lineFour"`  //
}

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

func (Dual150) New() Dual150 {
	return Dual150{
		Countdown: Line{}.SetDefault(),
		LineOne:   Line{}.SetDefault(),
		LineTwo:   Line{}.SetDefault(),
		LineThree: Line{}.SetDefault(),
		LineFour:  Line{}.SetDefault(),
	}
}

func (Dual150) ParseRawData(rawData string) (Dual150, error) {
	if len(rawData) < 2 {
		return Dual150{}, errors.New("invalid rawData input")
	}

	rawDataSplit := strings.Split(rawData, ":")

	var countdown Line

	lineOne, err := Line{}.Parse(rawDataSplit[3], rawDataSplit[4])
	if err != nil {
		return Dual150{}, err
	}

	lineTwo, err := Line{}.Parse(rawDataSplit[5], rawDataSplit[6])
	if err != nil {
		return Dual150{}, err
	}

	lineThree, err := Line{}.Parse(rawDataSplit[7], rawDataSplit[8])
	if err != nil {
		return Dual150{}, err
	}

	lineFour, err := Line{}.Parse(rawDataSplit[9], rawDataSplit[10])
	if err != nil {
		return Dual150{}, err
	}

	if !lineOne.IsZero() || !lineTwo.IsZero() || !lineThree.IsZero() || !lineFour.IsZero() {
		countdown = Line{}.SetDefault()
	} else {
		countdown, err = Line{}.CountdownParse(rawDataSplit[1])
		if err != nil {
			return Dual150{}, err
		}

		lineOne = Line{}.SetDefault()
		lineTwo = Line{}.SetDefault()
		lineThree = Line{}.SetDefault()
		lineFour = Line{}.SetDefault()
	}

	d := Dual150{
		Countdown: countdown,
		LineOne:   lineOne,
		LineTwo:   lineTwo,
		LineThree: lineThree,
		LineFour:  lineFour,
	}

	return d, nil
}
