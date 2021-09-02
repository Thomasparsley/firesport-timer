package kocab

import "time"

type Line struct {
	Time   time.Time `json:"time"`
	Status Status    `json:"status"`
}

func (Line) Parse(rawTime string, rawID string) (Line, error) {
	lineTime, err := parseRawTimeToTime(rawTime)
	if err != nil {
		return Line{}, err
	}

	lineStatus, err := Status{}.Parse(rawID)
	if err != nil {
		return Line{}, err
	}

	return Line{
		Time:   lineTime,
		Status: lineStatus,
	}, nil
}

func (Line) CountdownParse(rawTime string) (Line, error) {
	countdownTime, err := parseRawTimeToTime(rawTime)
	if err != nil {
		return Line{}, err
	}

	result := Line{
		Time:   countdownTime,
		Status: Status{}.GetByID(StatusUndefinedID),
	}

	if result.IsZero() {
		result.Status = Status{}.GetByID(StatusStopID)
	} else {
		result.Status = Status{}.GetByID(StatusRunID)
	}

	return result, nil
}

func (Line) SetDefault() Line {
	return Line{
		Time:   time.Date(1679, 1, 1, 0, 0, 0, 0, time.Local),
		Status: Status{}.GetByID(StatusDefaultID),
	}
}

func (l *Line) IsZero() bool {
	return l.Time.Minute() == 0 && l.Time.Second() == 0 && l.Time.Nanosecond() == 0
}
