package line

import (
	"fmt"
	"strconv"
	"time"

	"thomasparsley.cz/firesport-timer/timers/status"
)

type Line struct {
	Time   time.Time     `json:"time"`
	Status status.Status `json:"status"`
}

func (l Line) IsZero() bool {
	return l.Time.Minute() == 0 &&
		l.Time.Second() == 0 &&
		l.Time.Nanosecond() == 0
}
func (l *Line) SetDefault() {
	newLine := New()
	l = &newLine
}

func New() Line {
	return Line{
		Time:   time.Date(1, 1, 1, 0, 0, 0, 0, time.Local),
		Status: status.New(),
	}
}

func Parse(rawTime string, rawID string) (Line, error) {
	timeInt, err := strconv.Atoi(rawTime)
	if err != nil {
		return Line{}, err
	}

	t := time.Date(1, 1, 1, 0, 0, 0, timeInt*1000000, time.Local)

	s, err := status.ParseRaw(rawID)
	if err != nil {
		return Line{}, err
	}

	return Line{
		Time:   t,
		Status: s,
	}, nil
}

func ParseCountdown(rawTime string) (Line, error) {
	result, err := Parse(rawTime, fmt.Sprintf("%d", status.UndefinedID))
	if err != nil {
		return Line{}, err
	}

	if result.IsZero() {
		result.Status = status.GetByID(status.StopID)
	} else {
		result.Status = status.GetByID(status.RunID)
	}

	return result, nil
}
