package kocab

import (
	"strconv"
	"time"
)

func FormatTime(input time.Time) string {
	t := input.Format("4:05")

	// Add to time milliseconds
	milliseconds := strconv.Itoa(int(input.Nanosecond()))
	for len(milliseconds) < 9 {
		milliseconds = "0" + milliseconds
	}

	t += "." + milliseconds[:3]

	return t
}
