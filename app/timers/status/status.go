package status

import "strconv"

const (
	UndefinedID   = int(0)
	UndefinedName = string("undefined")
	DefaultID     = int(1)
	DefaultName   = string("default")
	RunID         = int(2)
	RunName       = string("run")
	StopID        = int(8)
	StopName      = string("stop")
)

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func New() Status {
	return newStatus(UndefinedID, UndefinedName)
}

func newStatus(id int, name string) Status {
	return Status{
		Id:   id,
		Name: name,
	}
}

func GetByID(id int) Status {
	switch id {
	case DefaultID:
		return newStatus(DefaultID, DefaultName)
	case RunID:
		return newStatus(RunID, RunName)
	case StopID:
		return newStatus(StopID, StopName)
	default:
		return newStatus(UndefinedID, UndefinedName)
	}
}

func ParseRaw(raw string) (Status, error) {
	id, err := strconv.Atoi(raw)
	if err != nil {
		return Status{}, err
	}

	return GetByID(id), nil
}
