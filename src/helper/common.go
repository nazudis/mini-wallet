package helper

import (
	"math/rand"
)

// default and minimum length is 8
func RandString(length ...int) string {
	n := 8
	if len(length) > 0 && length[0] >= 8 {
		n = length[0]
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomString := make([]byte, n)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomString)
}

func VarToPointer[T any](p T) *T {
	return &p
}
