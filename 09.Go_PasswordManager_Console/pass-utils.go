package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// Get user password from Terminal
func readPassword(prompt string) ([]byte, error) {
	fmt.Print(prompt)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println() // Print a newline after user input
	return password, err
}

// Provide user with password generation choice
func getPasswordChoice() ([]byte, error) {
	choice := 0
	var keyPassword []byte
	fmt.Printf("Make a choice:\n")
	fmt.Printf("(1)Auto-generate a password / (2) Provide your own password [1/2]:")
	fmt.Scanln(&choice)

	if choice == 1 {
		passLength := 0
		fmt.Println("Specify password length:")
		fmt.Scanln(&passLength)

		fmt.Println("Generating random password...")
		securePassword, err := GenerateSecurePassword(passLength)
		if err != nil {
			fmt.Println(err)
		}

		keyPassword = []byte(securePassword)

	} else {
		// Prompt user for their self provided key password
		fmt.Print("Enter the key password: ")
		keyPassword, _ = terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println() // Print a newline after user input
	}

	return keyPassword, nil
}

// Secure Password Generation
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
