package utils

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateSKU() (string, error) {
	date := time.Now().Format("20060102")

	randomPart, err := randomString(5)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("SKU-%s-%s", date, randomPart), nil
}

func randomString(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := range bytes {
		bytes[i] = charset[bytes[i]%byte(len(charset))]
	}

	return string(bytes), nil
}
