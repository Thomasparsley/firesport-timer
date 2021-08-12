package kocab

import "strconv"

const (
	StatusUndefinedID = int(0)
	StatusDefaultID   = int(1)
	StatusRunID       = int(2)
	StatusStopID      = int(8)
)

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (Status) GetByID(id int) Status {
	switch id {
	case 1:
		return Status{
			Id:   StatusDefaultID,
			Name: "default",
		}
	case 2:
		return Status{
			Id:   StatusRunID,
			Name: "run",
		}
	case 8:
		return Status{
			Id:   StatusStopID,
			Name: "stop",
		}
	default:
		return Status{
			Id:   StatusUndefinedID,
			Name: "undefined",
		}
	}
}

func (Status) Parse(rawID string) (Status, error) {
	id, err := strconv.Atoi(rawID)

	if err != nil {
		return Status{}, err
	}

	return Status{}.GetByID(id), nil
}
