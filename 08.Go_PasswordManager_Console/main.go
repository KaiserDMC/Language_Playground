package main

import (
	"bufio"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type Profile struct {
	Name          string
	CipherKey     string
	PasswordStash []Password
}

type Password struct {
	Website string
	Url     string
	Pass    string
}

func main() {
	var profile Profile
	var cipherKey, filePath string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Would you like to open an existing profile? [Y/n]:")
	input, err := reader.ReadString(`\n`)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	if input[0] == "Y" || input[0] == "y" {
		fmt.Println("Please provide a file path:")
		fmt.Scanln(&filePath)

		// Read existing profile
		existingProfile, err := readProfile(filePath)
		if err != nil {
			fmt.Println("Error reading profile:", err)
			return
		}
		profile = existingProfile

		fmt.Println("Please enter your Master Key (input is not visible):")
		bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading master key:", err)
			return
		}
		profile.CipherKey = string(bytePassword)
	} else {
		// Create a new profile
		fmt.Print("Enter the profile name: ")
		fmt.Scanln(&profile.Name)

		fmt.Print("Please enter your Master Key (input is not visible): ")
		bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading master key:", err)
			return
		}
		profile.CipherKey = string(bytePassword)
	}

	MainMenu(filePath, &profile)

}

func readProfile(filePath string) (Profile, error) {
	var profile Profile

	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return profile, err
	}

	// Decrypt the content using the cipher key
	decryptedContent := DecryptAES([]byte(profile.CipherKey), string(content))

	// Unmarshal the JSON into the profile struct
	err = json.Unmarshal([]byte(decryptedContent), &profile)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func writeProfile(filePath string, profile Profile) error {
	// Marshal the profile struct into JSON
	profileJSON, err := json.Marshal(profile)
	if err != nil {
		return err
	}

	// Encrypt the content using the cipher key
	encryptedContent := EncryptAES([]byte(profile.CipherKey), string(profileJSON))

	// Write the encrypted content to the file
	err = ioutil.WriteFile(filePath, []byte(encryptedContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func EncryptAES(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])

	fmt.Println(s)
}
