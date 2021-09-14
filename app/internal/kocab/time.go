package kocab

import (
	"strconv"
	"time"
)

func FormatTime(input time.Time) string {
	t := input.Format("4:05")

	// Add to time mili-seconds
	mili := strconv.Itoa(int(input.Nanosecond()))
	for len(mili) < 9 {
		mili = "0" + mili
	}

	t += "." + mili[:3]

	return t
}

func parseRawTimeToTime(rawTime string) (time.Time, error) {
	timeInt, err := strconv.Atoi(rawTime)
	if err != nil {
		return time.Time{}, err
	}

	t := time.Date(1679, 1, 1, 0, 0, 0, timeInt*1000000, time.Local)

	return t, nil
}
