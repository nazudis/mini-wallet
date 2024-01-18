package helper

import (
	"encoding/hex"
)

func GenToken(s string) string {
	return hex.EncodeToString([]byte(s))
}

func DecodeToken(s string) string {
	ds, _ := hex.DecodeString(s)
	return string(ds)
}
