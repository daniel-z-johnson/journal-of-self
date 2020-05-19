package random

import (
	"crypto/rand"
	"encoding/base64"
)

func nBytes(n int) ([]byte, error) {
	rBytes := make([]byte, n)
	_, err := rand.Read(rBytes)
	if err != nil {
		return nil, err
	}
	return rBytes, nil
}

// GenerateToken - Generates a n-byte token
// and returns it as a base64 string
func GenerateToken(numBytes int) (string, error) {
	bytes, err := nBytes(numBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil

}
