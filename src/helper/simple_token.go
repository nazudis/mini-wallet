package helper

import (
	"encoding/hex"
	"strings"
)

var salt = RandString(16)

func GenToken(s string) string {
	s = s + ":" + salt
	return hex.EncodeToString([]byte(s))
}

func DecodeToken(s string) string {
	ds, _ := hex.DecodeString(s)
	result := strings.Split(string(ds), ":")[0]
	return result
}
