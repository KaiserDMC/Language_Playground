package main

import (
	"fmt"
	"os"
)

func MainMenu(filePath string, profile *Profile) {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("[1] Create a new password")
		fmt.Println("[2] Look for a password inside current profile")
		fmt.Println("[0] Exit")
		fmt.Println("Choose an option:")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading choice:", err)
			break
		}

		switch choice {
		case 0:
			err := writeProfile(filePath, *profile)
			if err != nil {
				fmt.Println("Error writing profile:", err)

			}
			fmt.Println("Exiting...")
			os.Exit(0)
		case 1:
			// Create a new password
			var newPassword Password
			fmt.Print("Enter website name of the new password: ")
			fmt.Scanln(&newPassword.Website)
			fmt.Print("Enter URL for the new password: ")
			fmt.Scanln(&newPassword.Url)
			fmt.Print("Enter password for the new password: ")
			fmt.Scanln(&newPassword.Pass)

			profile.PasswordStash + append(profile.PasswordStash, newPassword)
			fmt.Println("Password added successfully!")
		case 2:
			// Look for a password by website
			var searchWebsite string
			fmt.Print("Enter website to search for: ")
			fmt.Scanln(&searchWebsite)

			// Search for the password in the profile
			found := false
			for _, password := range profile.PasswordStash {
				if password.Website == searchWebsite {
					fmt.Printf("Password found:\nWebsite: %s\nURL: %s\nPassword: %s\n", password.Website, password.Url, password.Pass)
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Password not found.")
			}
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}

}
