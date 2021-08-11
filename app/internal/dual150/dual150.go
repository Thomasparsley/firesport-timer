package dual150

import (
	"encoding/hex"
)

// Decode hex string to UTF-8 string
func DecodeHexString(s string) string {
	byteString, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return string(byteString)[2:]
}

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetStatus(s int) Status {
	switch s {
	case 1:
		return Status{
			Id:   1,
			Name: "default",
		}
	case 2:
		return Status{
			Id:   2,
			Name: "run",
		}
	case 8:
		return Status{
			Id:   8,
			Name: "stop",
		}
	default:
		return Status{
			Id:   0,
			Name: "undefined",
		}
	}
}
