package main

import "fmt"

func main() {
	// Read input, number of tabs n and start salary
	var n, salary int
	fmt.Scan(&n)
	fmt.Scan(&salary)

	var website string
	// Loop through each tab
	for i := 0; i < n; i++ {
		// Get name of currenly opened tab
		fmt.Scan(&website)

		// Check for unauthorized websites
		switch website {
		case "Facebook":
			salary -= 150
		case "Instagram":
			salary -= 100
		case "Reddit":
			salary -= 50
		}

		// If we lost all of our salary
		if salary <= 0 {
			fmt.Println("You have lost your salary.")
			return
		}
	}

	// If we still have any salary left
	fmt.Println(salary)
}
