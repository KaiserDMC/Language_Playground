package main

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/scrypt"
)

type PasswordProfile struct {
	Id       string
	Name     string
	Url      string
	Password string
}

const (
	keyLen       = 32 // Length of the derived key
	scryptN      = 16384
	scryptR      = 8
	scryptP      = 1
	scryptSalt   = "g0lang_b3st_lang"
	scryptParams = scryptN | scryptR<<16 | scryptP<<24
	configFile   = "config.txt"
)

func main() {
	// Create command flags
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	storageCmd := flag.NewFlagSet("storage", flag.ExitOnError)
	keyCmd := flag.NewFlagSet("key", flag.ExitOnError)

	storageName := storageCmd.String("name", "", "Specify the storage name")
	vaultPassword := storageCmd.String("vaultPassword", "", "Specify the vault's master password")

	keyStorage := keyCmd.String("storage", "", "Specify the storage name for the key")
	keyVaultPassword := keyCmd.String("vaultPassword", "", "Specify the vault's master password")
	keyName := keyCmd.String("name", "", "Specify the name for the key")
	keyUrl := keyCmd.String("url", "", "Specify the URL for the key")
	keyPassword := keyCmd.String("password", "", "Specify the password for the key")

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid command! In case you need more information you can use 'help' command!")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		if createCmd.Parsed() {
			switch createCmd.Arg(0) {
			case "storage":
				storageCmd.Parse(os.Args[3:])
				if storageCmd.Parsed() {
					if *storageName == "" || *vaultPassword == "" {
						fmt.Println("Please provide storage name and a vault master password")
						os.Exit(1)
					}

					// Check if the storage already exists
					if _, err := os.Stat(fmt.Sprintf("%s.db", *storageName)); err == nil {
						fmt.Printf("Storage with name %s already exists!\n", *storageName)
						os.Exit(1)
					}

					// Derive key from the vault password using scrypt
					derivedKey, err := scrypt.Key([]byte(*vaultPassword), []byte(scryptSalt), scryptN, scryptR, scryptP, keyLen)
					if err != nil {
						fmt.Println("Error deriving key:", err)
						os.Exit(1)
					}

					// Save the derived key and other necessary parameters to a configuration file
					configData := fmt.Sprintf("Key: %s\nSalt: %s\nN: %d\nR: %d\nP: %d\n", base64.StdEncoding.EncodeToString(derivedKey), scryptSalt, scryptN, scryptR, scryptP)
					err = ioutil.WriteFile(configFile, []byte(configData), 0600)
					if err != nil {
						fmt.Println("Error saving configuration:", err)
						os.Exit(1)
					}

					fmt.Printf("Creating storage with name: %s\n", *storageName)

					// Create sqlite database for storage of passwords
					db, err := sql.Open("sqlite3", fmt.Sprintf("%s.db", *storageName))
					if err != nil {
						fmt.Println("Error creating storage:", err)
						return
					}
					defer db.Close()

					fmt.Printf("%s password storage created!\n", *storageName)

					// Create Table
					_, err = db.Exec(`
									CREATE TABLE IF NOT EXISTS passwords (
										id TEXT PRIMARY KEY,
										name TEXT,
										url TEXT,
										pass TEXT
									)
								`)
					if err != nil {
						fmt.Println("Error creating table:", err)
						return
					}

					fmt.Println("Storage table created!")
				}

			case "key":
				keyCmd.Parse(os.Args[3:])
				if keyCmd.Parsed() {
					if *keyStorage == "" || *keyVaultPassword == "" {
						fmt.Println("Please provide storage name and vault master password")
						os.Exit(1)
					}

					// Load the stored configuration from the configuration file
					configData, err := ioutil.ReadFile(configFile)
					if err != nil {
						fmt.Println("Error reading configuration file:", err)
						os.Exit(1)
					}

					// Parse the configuration data
					var storedKey []byte
					var storedSalt string
					var storedN, storedR, storedP int
					_, err = fmt.Sscanf(string(configData), "Key: %s\nSalt: %s\nN: %d\nR: %d\nP: %d\n",
						&storedKey, &storedSalt, &storedN, &storedR, &storedP)
					if err != nil {
						fmt.Println("Error parsing configuration data:", err)
						os.Exit(1)
					}

					// Derive key from the user-provided vault password
					userDerivedKey, err := scrypt.Key([]byte(*keyVaultPassword), []byte(scryptSalt), scryptN, scryptR, scryptP, keyLen)
					if err != nil {
						fmt.Println("Error deriving key:", err)
						os.Exit(1)
					}

					fmt.Printf("Stored Key: %x\n", storedKey)
					fmt.Printf("User Derived Key: %x\n", userDerivedKey)

					// Compare the derived keys
					if !bytes.Equal(userDerivedKey, storedKey) {
						fmt.Println("Vault password does not match. Access denied.")
						os.Exit(1)
					}

					// Derive key from the vault password using scrypt
					derivedKey, err := scrypt.Key([]byte(*keyVaultPassword), []byte(scryptSalt), scryptN, scryptR, scryptP, keyLen)
					if err != nil {
						fmt.Println("Error deriving key:", err)
						os.Exit(1)
					}

					// Encrypt the password using the derived key
					encryptedPassword, err := encrypt([]byte(*keyPassword), derivedKey)
					if err != nil {
						fmt.Println("Error encrypting password:", err)
						os.Exit(1)
					}

					// Create PasswordProfile and insert it into the storage
					passwordProfile := PasswordProfile{
						Id:       uuid.New().String(),
						Name:     *keyName,
						Url:      *keyUrl,
						Password: string(encryptedPassword),
					}

					// Use the storage name to open the appropriate database
					db, err := sql.Open("sqlite3", fmt.Sprintf("%s.db", *keyStorage))
					if err != nil {
						fmt.Println("Error opening database:", err)
						os.Exit(1)
					}
					defer db.Close()

					// Insert the PasswordProfile into the storage
					_, err = db.Exec(`
						INSERT INTO passwords (id, name, url, pass)
						VALUES (?, ?, ?, ?)
					`, passwordProfile.Id, passwordProfile.Name, passwordProfile.Url, passwordProfile.Password)
					if err != nil {
						fmt.Println("Error inserting data:", err)
						os.Exit(1)
					}

					fmt.Println("PasswordProfile inserted into storage successfully!")
				}
			default:
				fmt.Println("Unknown sub-command under 'create'. Please use 'storage'")
				os.Exit(1)
			}
		}
	case "help":
		helpCmd.Parse(os.Args[2:])
		// Handle 'help' command logic here
		fmt.Println("Displaying help...")
	default:
		fmt.Println("Unknown command. Please use 'create' or 'help'")
		os.Exit(1)
	}
}

// Encrypt encrypts the data using the provided key.
func encrypt(data, key []byte) ([]byte, error) {
	// You need to implement encryption logic here.
	// This could involve using a symmetric encryption algorithm like AES.
	// For simplicity, you can use a library like golang.org/x/crypto/chacha20poly1305.
	// Replace the following line with your actual encryption logic.
	return data, nil
}
