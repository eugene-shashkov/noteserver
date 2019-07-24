package utils

import (
	"crypto/rand"
	"fmt"
)

// RandomBytes is a func to generate random byte string. rand.Read returns random bytes in the range 0-255
func RandomBytes() string {
	a := make([]byte, 8)
	rand.Read(a)
	return fmt.Sprintf("%x", a)
}

// GenerateToken is a  token generator
func GenerateToken() string {
	return RandomBytes() + RandomBytes() + RandomBytes() + RandomBytes()
}
