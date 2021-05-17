package bufferTransfare

import "encoding/hex"

// DecodeHexString
// decode hex string to UTF-8 string
func DecodeHexString(s string) string {
	byteString, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return string(byteString)
}
