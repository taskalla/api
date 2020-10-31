package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateToken() string {
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
