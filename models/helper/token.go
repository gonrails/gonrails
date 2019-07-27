package helper

import (
	"crypto/rand"
	"fmt"
)

// GenerateToken - for Models token
func GenerateToken() string {
	defaultLen := 16
	return generateToken(defaultLen)
}

func generateToken(len int) string {
	bytes := make([]byte, len)
	rand.Read(bytes)
	return fmt.Sprintf("%x", bytes)
}

// RSA tools
