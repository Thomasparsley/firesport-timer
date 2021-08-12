package kocab

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

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

func parseRawTimeToTime(rawTime string) (time.Time, error) {
	timeInt, err := strconv.Atoi(rawTime)
	if err != nil {
		return time.Time{}, err
	}

	t := time.Date(1, 1, 1, 0, 0, 0, timeInt*1000000, time.Local)

	return t, nil
}

func ParseRawData(rawData string) (dual150, error) {
	if len(rawData) < 2 {
		return dual150{}, errors.New("invalid rawData input")
	}

	rawDataSplit := strings.Split(rawData, ":")

	countdown, err := Line{}.CountdownParse(rawDataSplit[1])
	if err != nil {
		return dual150{}, err
	}

	// left line
	lineOne, err := Line{}.Parse(rawDataSplit[3], rawDataSplit[4])
	if err != nil {
		return dual150{}, err
	}

	// right line
	lineTwo, err := Line{}.Parse(rawDataSplit[5], rawDataSplit[6])
	if err != nil {
		return dual150{}, err
	}

	d := dual150{
		Countdown: countdown,
		LineOne:   lineOne,
		LineTwo:   lineTwo,
		LineThree: Line{}.SetDefault(),
		LineFour:  Line{}.SetDefault(),
	}

	return d, nil
}
