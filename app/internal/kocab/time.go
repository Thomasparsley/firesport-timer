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
