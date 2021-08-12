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

	var countdownStatus Status
	if countdownTime.Second() == 0 && countdownTime.Minute() == 0 {
		countdownStatus = Status{}.GetByID(StatusStopID)
	} else {
		countdownStatus = Status{}.GetByID(StatusRunID)
	}

	return Line{
		Status: countdownStatus,
		Time:   countdownTime,
	}, nil
}

func (Line) SetDefault() Line {
	return Line{
		Time:   time.Date(1, 1, 1, 0, 0, 0, 0*1000000, time.Local),
		Status: Status{}.GetByID(StatusDefaultID),
	}
}
