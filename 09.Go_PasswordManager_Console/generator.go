package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
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

// Extracted function to handle password choice and generation
func getPasswordChoice() ([]byte, error) {
	choice := 0
	var keyPassword []byte

	fmt.Print("Make a choice:\n(1) Auto-generate a password / (2) Provide your own password [1/2]: ")

	bufio.NewReader(os.Stdin).ReadBytes('\n')

	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Error reading choice:", err)
		return nil, err
	}

	if choice == 1 {
		passLength := 0
		fmt.Print("Specify password length: ")
		_, err := fmt.Scanln(&passLength)
		if err != nil {
			fmt.Println("Error reading password length:", err)
			return nil, err
		}

		fmt.Println("Generating random password...")
		securePassword, err := GenerateSecurePassword(passLength)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		keyPassword = []byte(securePassword)

	} else {
		// Prompt user for the key password
		fmt.Print("Enter the key password: ")
		keyPassword, err = terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println() // Print a newline after user input
		if err != nil {
			fmt.Println("Error reading key password:", err)
			return nil, err
		}
	}

	return keyPassword, nil
}
