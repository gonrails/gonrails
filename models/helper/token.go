package helper

import (
	"crypto/rand"
	"fmt"
)

func GenerateToken() string {
	defaultLen := 16
	bytes := make([]byte, defaultLen)
	rand.Read(bytes)
	return fmt.Sprintf("%x", bytes)
}
