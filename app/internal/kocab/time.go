package kocab

import (
	"strconv"
	"time"
)

func FormatTime(input time.Time) string {
	t := input.Format("04:05")

	// Add to time mili-seconds
	mili := strconv.Itoa(int(input.Nanosecond()))
	for len(mili) < 3 {
		mili = "0" + mili
	}

	t += "." + mili[:3]

	return t
}
