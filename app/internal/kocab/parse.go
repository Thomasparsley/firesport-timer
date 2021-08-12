package kocab

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

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
	Odpočet
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
