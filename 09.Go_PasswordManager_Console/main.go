package main

import (
	"database/sql"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/ssh/terminal"
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

func main() {
	// Create command flags
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	vaultCmd := flag.NewFlagSet("vault", flag.ExitOnError)
	keyCmd := flag.NewFlagSet("key", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	cleanCmd := flag.NewFlagSet("clean", flag.ExitOnError)

	vaultName := vaultCmd.String("name", "", "Specify the vault name")

	keyVault := keyCmd.String("vault", "", "Specify the vault name for the key")
	keyName := keyCmd.String("name", "", "Specify the name for the key")
	keyUrl := keyCmd.String("url", "", "Specify the URL for the key")

	listVaults := listCmd.Bool("vaults", false, "List all vaults")
	listKeys := listCmd.Bool("keys", false, "List keys in a vault")
	listVaultName := listCmd.String("vaultName", "", "Specify the vault name for listing keys")

	showVaultName := showCmd.String("vaultName", "", "Specify the vault name for showing a key")
	showKeyName := showCmd.String("key", "", "Specify the key name for showing")

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	// Extract the provided by the user SALT
	scryptSalt, err := os.ReadFile("user-salt.txt")
	if err != nil {
		fmt.Println("Error reading 'user-salt' file:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid command! In case you need more information you can use 'help' command!")
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
					fmt.Print("Enter the vault password: ")
					vaultPassword, err := terminal.ReadPassword(int(syscall.Stdin))
					fmt.Println() // Print a newline after user input
					if err != nil {
						fmt.Println("Error reading vault password:", err)
						os.Exit(1)
					}

					// Check if "vaults" folder exists
					folderErr := CreateFolderIfNotExists("vaults")
					if folderErr != nil {
						fmt.Println("Error creating 'vaults' folder:", folderErr)
						return
					}

					// Check if the vault already exists
					vaultDBPath := fmt.Sprintf("./vaults/%s.db", *vaultName)
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
					db, err := sql.Open("sqlite3", fmt.Sprintf("./vaults/%s.db", *vaultName))
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
				}

			case "key":
				keyCmd.Parse(os.Args[3:])
				if keyCmd.Parsed() {
					if *keyVault == "" || *keyName == "" || *keyUrl == "" {
						fmt.Println("Please provide vault name, key name, and key URL")
						os.Exit(1)
					}

					// Prompt user for the vault password
					fmt.Print("Enter the vault password: ")
					vaultPassword, err := terminal.ReadPassword(int(syscall.Stdin))
					fmt.Println() // Print a newline after user input
					if err != nil {
						fmt.Println("Error reading vault password:", err)
						os.Exit(1)
					}

					// Prompt user for the key password
					fmt.Print("Enter the key password: ")
					keyPassword, err := terminal.ReadPassword(int(syscall.Stdin))
					fmt.Println() // Print a newline after user input
					if err != nil {
						fmt.Println("Error reading key password:", err)
						os.Exit(1)
					}

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

					//fmt.Printf("Stored Key: %x\n", decodedKey)

					// Derive key from the vault password using scrypt
					derivedKey, err := DeriveUserKey(string(vaultPassword))
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					// Print the derived key as a hexadecimal string
					//fmt.Printf("User Derived Key: %x\n", derivedKey)

					// Compare the derived keys
					err = CompareKeys(decodedKey, derivedKey)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					//fmt.Printf("User Derived Key: %x\n", derivedKey)

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
					if _, err := os.Stat("vaults"); os.IsNotExist(err) {
						if err != nil {
							fmt.Println("Error 'vaults' folder does not exist", err)
							return
						}
					}

					// Use the vault name to open the appropriate database
					db, err := sql.Open("sqlite3", fmt.Sprintf("./vaults/%s.db", *keyVault))
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

	case "help":
		helpCmd.Parse(os.Args[2:])
		// Handle 'help' command logic here
		fmt.Println("Displaying help...")

	case "list":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a valid sub-command for listing: 'vaults' or 'keys'")
			os.Exit(1)
		}

		listCmd.Parse(os.Args[2:])

		if *listVaults {
			// List all vaults
			files, err := os.ReadDir("vaults")
			if err != nil {
				fmt.Println("Error listing vaults:", err)
				os.Exit(1)
			}

			fmt.Println("List of Vaults:")
			for _, file := range files {
				fmt.Println(file.Name())
			}
		} else if *listKeys {
			// List keys in a vault
			if *listVaultName == "" {
				fmt.Println("Please provide a vault name using '-vaultName' flag for listing keys")
				os.Exit(1)
			}

			vaultName := *listVaultName

			// Prompt user for the vault password
			fmt.Print("Enter the vault password: ")
			vaultPassword, err := terminal.ReadPassword(int(syscall.Stdin))
			fmt.Println() // Print a newline after user input
			if err != nil {
				fmt.Println("Error reading vault password:", err)
				os.Exit(1)
			}

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

			//fmt.Printf("Stored Key: %x\n", decodedKey)

			// Derive key from the vault password using scrypt
			derivedKey, err := DeriveUserKey(string(vaultPassword))
			if err != nil {
				fmt.Println("Error deriving key:", err)
				os.Exit(1)
			}

			// Print the derived key as a hexadecimal string
			//fmt.Printf("User Derived Key: %x\n", derivedKey)

			// Compare the derived keys
			err = CompareKeys(decodedKey, derivedKey)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			//fmt.Printf("User Derived Key: %x\n", derivedKey)

			// Use the vault name to open the appropriate database
			db, err := sql.Open("sqlite3", fmt.Sprintf("./vaults/%s.db", vaultName))
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

		} else {
			fmt.Println("Please provide a valid sub-command for 'list': '-vaults' or '-keys'")
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
		fmt.Print("Enter the vault password: ")
		vaultPassword, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println() // Print a newline after user input
		if err != nil {
			fmt.Println("Error reading vault password:", err)
			os.Exit(1)
		}

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

		// Use the vault name to open the appropriate database
		db, err := sql.Open("sqlite3", fmt.Sprintf("./vaults/%s.db", vaultName))
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

		// Print the decrypted password to the terminal
		//fmt.Printf("Decrypted Password for Key %s in Vault %s: %s\n", keyName, vaultName, decryptedPassword)

		// Copy the decrypted password to the clipboard
		err = clipboard.WriteAll(string(decryptedPassword))
		if err != nil {
			fmt.Println("Error copying password to clipboard:", err)
			os.Exit(1)
		}

		fmt.Println("Password copied to clipboard.")

	case "clean":
		cleanCmd.Parse(os.Args[2:])
		if cleanCmd.Parsed() {
			err := cleanUnencryptedFiles(cryptConfigs)
			if err != nil {
				fmt.Println("Error cleaning unencrypted files:", err)
				os.Exit(1)
			}
			fmt.Println("Unencrypted files cleaned successfully.")
		}
	default:
		fmt.Println("Unknown command. Please use 'create' or 'help'")
		os.Exit(1)
	}
}
