package main

import (
	"database/sql"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
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
	scryptSalt   = ""
	scryptParams = scryptN | scryptR<<16 | scryptP<<24
	cryptConfigs = "crypt-configs"
	vaultFolder  = "vaults"
)

// Create command flags
var (
	createCmd = flag.NewFlagSet("create", flag.ExitOnError)
	vaultCmd  = flag.NewFlagSet("vault", flag.ExitOnError)
	keyCmd    = flag.NewFlagSet("key", flag.ExitOnError)
	showCmd   = flag.NewFlagSet("show", flag.ExitOnError)

	// Flags for vault creation cmd
	vaultName = vaultCmd.String("name", "", "Specify the vault name")

	// Flags for key creation cmd
	keyVault = keyCmd.String("vaultName", "", "Specify the vault name for the key")
	keyName  = keyCmd.String("key", "", "Specify the name for the key")
	keyUrl   = keyCmd.String("url", "", "Specify the URL for the key")

	// Flags for show cmd
	showVaultName = showCmd.String("vaultName", "", "Specify the vault name for showing a key")
	showKeyName   = showCmd.String("key", "", "Specify the key name for showing")
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid command! In case you need more information you can use 'help' command!")
		os.Exit(1)
	}

	// Extract the provided by the user SALT
	scryptSalt, err := os.ReadFile(".user-salt")
	if err != nil {
		fmt.Println("Error reading 'user-salt' file:", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		if createCmd.Parsed() {
			switch createCmd.Arg(0) {
			case "vault":
				vaultCmd.Parse(os.Args[3:])
				if vaultCmd.Parsed() {
					if *vaultName == "" {
						fmt.Println("Please provide a vault name")
						os.Exit(1)
					}

					// Prompt user for the vault password
					vaultPassword, err := readPassword("Enter the vault password: ")
					if err != nil {
						fmt.Println("Error reading vault password:", err)
						os.Exit(1)
					}

					// Check if "vaults" folder exists
					folderErr := CreateFolderIfNotExists(vaultFolder)
					if folderErr != nil {
						fmt.Println("Error creating 'vaults' folder:", folderErr)
						return
					}

					// Check if the vault already exists
					vaultDBPath := fmt.Sprintf("./%s/%s.db", vaultFolder, *vaultName)
					if _, err := os.Stat(vaultDBPath); err == nil {
						fmt.Printf("Vault with name %s already exists!\n", *vaultName)
						os.Exit(1)
					}

					// Derive key from the vault password using scrypt
					derivedKey, err := DeriveUserKey(string(vaultPassword))
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					// Create a folder for storing encrypted passwords
					err = CreateFolderIfNotExists(cryptConfigs)
					if err != nil {
						fmt.Println("Error creating 'crypt-configs' folder:", err)
						os.Exit(1)
					}

					// Save the derived key and other necessary parameters to a configuration file
					configData := fmt.Sprintf("Key: %s\nSalt: %s\nN: %d\nR: %d\nP: %d\n", base64.StdEncoding.EncodeToString(derivedKey), scryptSalt, scryptN, scryptR, scryptP)
					err = SaveToFile(filepath.Join(cryptConfigs, *vaultName), []byte(configData))
					if err != nil {
						fmt.Println("Error saving configuration:", err)
						os.Exit(1)
					}

					// Encrypt and save the vault password to a file
					encryptedPassword, err := encrypt([]byte(vaultPassword), derivedKey)
					if err != nil {
						fmt.Println("Error encrypting password:", err)
						os.Exit(1)
					}

					// Save the encrypted password to a file named after the vault
					encryptedPasswordFile := filepath.Join(cryptConfigs, fmt.Sprintf("%s.enc", *vaultName))
					err = SaveToFile(encryptedPasswordFile, encryptedPassword)
					if err != nil {
						fmt.Println("Error saving encrypted password:", err)
						os.Exit(1)
					}

					fmt.Printf("Creating vault with name: %s\n", *vaultName)

					// Create sqlite database for storage of passwords
					db, err := sql.Open("sqlite3", fmt.Sprintf("./%s/%s.db", vaultFolder, *vaultName))
					if err != nil {
						fmt.Println("Error creating vault:", err)
						return
					}
					defer db.Close()

					fmt.Printf("%s password vault created!\n", *vaultName)

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

					hideUnencryptedFiles(cryptConfigs)
				}

			case "key":
				keyCmd.Parse(os.Args[3:])
				if keyCmd.Parsed() {
					if *keyVault == "" || *keyName == "" || *keyUrl == "" {
						fmt.Println("Please provide vault name, key name, and key URL")
						os.Exit(1)
					}

					// Prompt user for the vault password
					vaultPassword, err := readPassword("Enter the vault password: ")
					if err != nil {
						fmt.Println("Error reading vault password:", err)
						os.Exit(1)
					}

					// Password choice and generation
					keyPassword, err := getPasswordChoice()
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					showHiddenFiles(cryptConfigs)

					// Load the stored configuration from the configuration file
					configData, err := LoadConfig(*keyVault)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					// Parse the configuration data
					storedKey, err := ParseConfig(configData)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					// Decode the stored key from base64
					decodedKey, err := DecodeKey(storedKey)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					// Derive key from the vault password using scrypt
					derivedKey, err := DeriveUserKey(string(vaultPassword))
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					// Compare the derived keys
					err = CompareKeys(decodedKey, derivedKey)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					// Encrypt the password using the derived key
					encryptedPassword, err := encrypt([]byte(keyPassword), derivedKey)
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

					// Check if "vaults" folder exists
					if _, err := os.Stat(vaultFolder); os.IsNotExist(err) {
						if err != nil {
							fmt.Println("Error 'vaults' folder does not exist", err)
							return
						}
					}

					hideUnencryptedFiles(cryptConfigs)

					// Use the vault name to open the appropriate database
					db, err := sql.Open("sqlite3", fmt.Sprintf("./%s/%s.db", vaultFolder, *keyVault))
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
			}
		}

	case "list":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a valid sub-command for listing: 'vaults' or 'keys'")
			os.Exit(1)
		}

		subCommand := os.Args[2]
		switch subCommand {
		case "vaults":
			// List all vaults
			files, err := os.ReadDir(vaultFolder)
			if err != nil {
				fmt.Println("Error listing vaults:", err)
				os.Exit(1)
			}

			fmt.Println("List of Vaults:")
			for _, file := range files {
				fmt.Println(file.Name())
			}
		case "keys":
			if len(os.Args) < 4 {
				fmt.Println("Please provide a vault name using '-vaultName' flag for listing keys")
				os.Exit(1)
			}

			listVaultName := os.Args[4]
			vaultName := listVaultName

			// Prompt user for the vault password
			vaultPassword, err := readPassword("Enter the vault password: ")
			if err != nil {
				fmt.Println("Error reading vault password:", err)
				os.Exit(1)
			}

			showHiddenFiles(cryptConfigs)

			// Load the stored configuration from the configuration file
			configData, err := LoadConfig(vaultName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Parse the configuration data
			storedKey, err := ParseConfig(configData)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Decode the stored key from base64
			decodedKey, err := DecodeKey(storedKey)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Derive key from the vault password using scrypt
			derivedKey, err := DeriveUserKey(string(vaultPassword))
			if err != nil {
				fmt.Println("Error deriving key:", err)
				os.Exit(1)
			}

			// Compare the derived keys
			err = CompareKeys(decodedKey, derivedKey)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			hideUnencryptedFiles(cryptConfigs)

			// Use the vault name to open the appropriate database
			db, err := sql.Open("sqlite3", fmt.Sprintf("./%s/%s.db", vaultFolder, vaultName))
			if err != nil {
				fmt.Println("Error opening database:", err)
				os.Exit(1)
			}
			defer db.Close()

			// Query the database for name and url columns
			rows, err := db.Query("SELECT name, url FROM passwords")
			if err != nil {
				fmt.Println("Error querying database:", err)
				os.Exit(1)
			}
			defer rows.Close()

			fmt.Printf("List of Keys in Vault %s:\n", vaultName)
			index := 1
			for rows.Next() {
				var name, url string
				if err := rows.Scan(&name, &url); err != nil {
					fmt.Println("Error scanning database rows:", err)
					os.Exit(1)
				}

				fmt.Printf("%v. Name: %s, URL: %s\n", index, name, url)
				index++
			}

		default:
			fmt.Println("Please provide a valid sub-command for 'list': 'vaults' or 'keys'")
			os.Exit(1)
		}
	case "show":
		showCmd.Parse(os.Args[2:])

		vaultName := *showVaultName
		keyName := *showKeyName

		// Check if required flags are provided
		if vaultName == "" || keyName == "" {
			fmt.Println("Please provide valid arguments for showing a key: '-vaultName' and '-key'")
			os.Exit(1)
		}

		// Prompt user for the vault password
		vaultPassword, err := readPassword("Enter the vault password: ")
		if err != nil {
			fmt.Println("Error reading vault password:", err)
			os.Exit(1)
		}

		showHiddenFiles(cryptConfigs)

		// Load the stored configuration from the configuration file
		configData, err := LoadConfig(vaultName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Parse the configuration data
		storedKey, err := ParseConfig(configData)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Decode the stored key from base64
		decodedKey, err := DecodeKey(storedKey)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Derive key from the vault password using scrypt
		derivedKey, err := DeriveUserKey(string(vaultPassword))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Compare the derived keys
		err = CompareKeys(decodedKey, derivedKey)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		hideUnencryptedFiles(cryptConfigs)

		// Use the vault name to open the appropriate database
		db, err := sql.Open("sqlite3", fmt.Sprintf("./%s/%s.db", vaultFolder, vaultName))
		if err != nil {
			fmt.Println("Error opening database:", err)
			os.Exit(1)
		}
		defer db.Close()

		// Query the database for the password of the specified key
		var password string
		err = db.QueryRow("SELECT pass FROM passwords WHERE name = ?", keyName).Scan(&password)
		switch {
		case err == sql.ErrNoRows:
			fmt.Printf("Key with name %s not found in vault %s.\n", keyName, vaultName)
			os.Exit(1)
		case err != nil:
			fmt.Println("Error querying database:", err)
			os.Exit(1)
		}

		// Decrypt the password using the derived key
		decryptedPassword, err := decrypt([]byte(password), derivedKey)
		if err != nil {
			fmt.Println("Error decrypting password:", err)
			os.Exit(1)
		}

		// Copy the decrypted password to the clipboard
		err = clipboard.WriteAll(string(decryptedPassword))
		if err != nil {
			fmt.Println("Error copying password to clipboard:", err)
			os.Exit(1)
		}

		fmt.Println("Password copied to clipboard.")

	case "hide":
		hideUnencryptedFiles(cryptConfigs)
		fmt.Println("Unencrypted files hidden successfully.")
		os.Exit(0)

	case "delete":
		if len(os.Args) < 4 || os.Args[2] != "-vaultName" {
			fmt.Println("Please provide a valid sub-command for deleting: 'delete -vaultName <vaultName>'")
			os.Exit(1)
		}

		vaultToDelete := os.Args[3]

		if err := deleteVault(vaultToDelete); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "help":
		printHelp()

		os.Exit(0)

	default:
		fmt.Println("Unknown command:", flag.Arg(0))
		os.Exit(1)
	}
}
