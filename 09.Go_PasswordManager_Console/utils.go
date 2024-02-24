package main

import (
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
