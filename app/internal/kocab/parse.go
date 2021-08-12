package kocab

import (
	"strconv"
	"time"
)

func parseRawTimeToTime(rawTime string) (time.Time, error) {
	timeInt, err := strconv.Atoi(rawTime)
	if err != nil {
		return time.Time{}, err
	}

	t := time.Date(1, 1, 1, 0, 0, 0, timeInt*1000000, time.Local)

	return t, nil
}
