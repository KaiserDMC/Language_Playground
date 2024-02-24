package main

import (
	"flag"
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

// Function to clean unencrypted files
func cleanUnencryptedFiles(dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".enc") {
			filePath := filepath.Join(dirPath, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Function to hide unencrypted files
func hideUnencryptedFiles(dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	hiddenDir := filepath.Join(dirPath, ".hidden")
	err = os.Mkdir(hiddenDir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".enc") {
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
	fmt.Println("  create vault -name <vault_name>")
	fmt.Println("  create key -vault <vault_name> -name <key_name> -url <key_url>")
	fmt.Println("  list -vaults")
	fmt.Println("  list -keys -vaultName <vault_name>")
	fmt.Println("  show -vaultName <vault_name> -key <key_name>")
	fmt.Println("  hide -vaultName <vault_name> -key <key_name>")

	fmt.Println("\nFlags:")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("  -%s\t%s\n", f.Name, f.Usage)
	})
}
