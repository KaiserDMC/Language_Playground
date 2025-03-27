package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

func LoadConfig(keyVault string) ([]byte, error) {
	// Load the stored configuration from the configuration file
	configData, err := os.ReadFile(filepath.Join(cryptConfigs, keyVault))
	if err != nil {
		return nil, fmt.Errorf("error reading configuration file: %w", err)
	}

	return configData, nil
}

func ParseConfig(configData []byte) (string, error) {
	// Parse the configuration data
	var storedKey string
	_, err := fmt.Sscanf(string(configData), "Key: %s\n",
		&storedKey)
	if err != nil {
		return "", fmt.Errorf("error parsing configuration data: %w", err)
	}

	return storedKey, nil
}

func DecodeKey(storedKey string) ([]byte, error) {
	// Decode the stored key from base64
	decodedKey, err := base64.StdEncoding.DecodeString(storedKey)
	if err != nil {
		return nil, fmt.Errorf("error decoding stored key: %w", err)
	}

	//fmt.Printf("Stored Key: %x\n", decodedKey)
	return decodedKey, nil
}

func DeriveUserKey(keyVaultPassword string) ([]byte, error) {
	// Derive key from the vault password using scrypt
	derivedKey, err := DeriveKey(keyVaultPassword)
	if err != nil {
		return nil, fmt.Errorf("error deriving key: %w", err)
	}

	return derivedKey, nil
}

func CompareKeys(decodedKey, derivedKey []byte) error {
	// Compare the derived keys
	if !bytes.Equal(decodedKey, derivedKey) {
		return fmt.Errorf("vault password does not match - access denied")
	}

	return nil
}
