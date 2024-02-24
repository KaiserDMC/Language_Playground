package main

import (
	"crypto/rand"
	"math/big"
)

func GenerateSecurePassword(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>/?`~"

	maxIndex := big.NewInt(int64(len(chars)))
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			return "", err
		}

		password[i] = chars[randomIndex.Int64()]
	}

	return string(password), nil
}
