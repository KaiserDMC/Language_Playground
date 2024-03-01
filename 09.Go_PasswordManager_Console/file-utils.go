package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CreateFolderIfNotExists creates a folder if it doesn't exist.
func CreateFolderIfNotExists(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return os.MkdirAll(folderPath, 0755)
	}
	return nil
}

// SaveToFile saves data to a file with the given filepath.
func SaveToFile(filepath string, data []byte) error {
	return os.WriteFile(filepath, data, 0600)
}

// Function to delete a specified vault
func deleteVault(vaultName string) error {
	// Construct file paths
	configFile := filepath.Join(cryptConfigs, vaultName)
	encryptedFile := configFile + ".enc"
	hiddenConfigFile := filepath.Join(cryptConfigs, ".hidden", vaultName)
	hiddenEncryptedFile := hiddenConfigFile + ".enc"
	dbFile := filepath.Join(vaultFolder, vaultName+".db")

	// Check and delete the regular config file
	if err := os.Remove(configFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error deleting regular config file: %v", err)
	}

	// Check and delete the encrypted config file
	if err := os.Remove(encryptedFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error deleting encrypted config file: %v", err)
	}

	// Check and delete the hidden config file
	if err := os.Remove(hiddenConfigFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error deleting hidden config file: %v", err)
	}

	// Check and delete the hidden encrypted config file
	if err := os.Remove(hiddenEncryptedFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error deleting hidden encrypted config file: %v", err)
	}

	// Check and delete the database file
	if err := os.Remove(dbFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error deleting database file: %v", err)
	}

	fmt.Printf("Vault '%s' and associated files deleted successfully.\n", vaultName)
	return nil
}

// Function to hide unencrypted files
func hideUnencryptedFiles(dirPath string) error {
	// Check if ".hidden" directory already exists
	hiddenDir := filepath.Join(dirPath, ".hidden")
	_, err := os.Stat(hiddenDir)

	// If ".hidden" directory doesn't exist, proceed to creating it
	if err != nil {
		hiddenDir := filepath.Join(dirPath, ".hidden")
		err = os.Mkdir(hiddenDir, 0755)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}

	// Move existing unencrypted files (excluding those ending with ".enc") to ".hidden"
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("error listing files: %v", err)
	}

	for _, file := range files {
		if file.Name() != ".hidden" && !strings.HasSuffix(file.Name(), ".enc") {
			srcPath := filepath.Join(dirPath, file.Name())
			destPath := filepath.Join(hiddenDir, file.Name())

			err := os.Rename(srcPath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Function to show hidden files
func showHiddenFiles(dirPath string) error {
	hiddenDir := filepath.Join(dirPath, ".hidden")
	files, err := os.ReadDir(hiddenDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		srcPath := filepath.Join(hiddenDir, file.Name())
		destPath := filepath.Join(dirPath, file.Name())

		err := os.Rename(srcPath, destPath)
		if err != nil {
			return err
		}
	}

	err = os.Remove(hiddenDir)
	if err != nil {
		return err
	}

	return nil
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Printf("%s <command> [flags]\n", os.Args[0])
	fmt.Println("\nCommands:")
	fmt.Println("\n[1.] Create a new Password Storage Vault:")
	fmt.Println("  create vault -vaultName <vault_name>")
	fmt.Println("\n[2.] Create a new Key inside of already existing Password Storage Vault:")
	fmt.Println("  create key -vaultName <vault_name> -key <key_name> -url <key_url>")
	fmt.Println("\n[3.] List all currently available Password Storage Vaults:")
	fmt.Println("  list vaults")
	fmt.Println("\n[4.] List the name and url of all Keys in a specific Password Storage Vault:")
	fmt.Println("  list keys -vaultName <vault_name>")
	fmt.Println("\n[5.] Copy the requested Key from Password Storage Vault to User's Clipboard:")
	fmt.Println("  show -vaultName <vault_name> -key <key_name>")
	fmt.Println("\n[6.] Hide sensitive data files:")
	fmt.Println("  hide")
	fmt.Println("\n[7.] Delete specific Password Storage Vault and related files:")
	fmt.Println("  delete -vaultName <vault_name>")
}
