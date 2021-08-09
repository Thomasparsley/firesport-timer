package dual150

import (
	"encoding/hex"
)

type status struct {
	id   int
	name string
}

// DecodeHexString
// decode hex string to UTF-8 string
func DecodeHexString(s string) string {
	byteString, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return string(byteString)[2:]
}

func GetStatus(s int) status {
	switch s {
	case 1:
		return status{
			id:   1,
			name: "default",
		}
	case 2:
		return status{
			id:   2,
			name: "run",
		}
	case 8:
		return status{
			id:   8,
			name: "stop",
		}
	default:
		return status{
			id:   0,
			name: "undefinet",
		}
	}
}
