package main

import "fmt"

func main() {
	// Read input, number of groups, for each group number of people
	var groups int
	fmt.Scan(&groups)

	var countMusala, countMontblanc, countKilimandjaro, countK2, countEverest int
	var totalAmountOfPeople int
	// For each group reach the number of people
	for i := 0; i < groups; i++ {
		var numberOfPeople int
		fmt.Scan(&numberOfPeople)

		// Check which mountain top they will climb and add the people to their group
		if numberOfPeople <= 5 {
			countMusala += numberOfPeople
		} else if numberOfPeople >= 6 && numberOfPeople <= 12 {
			countMontblanc += numberOfPeople
		} else if numberOfPeople >= 13 && numberOfPeople <= 25 {
			countKilimandjaro += numberOfPeople
		} else if numberOfPeople >= 26 && numberOfPeople <= 40 {
			countK2 += numberOfPeople
		} else {
			countEverest += numberOfPeople
		}

		// Number of total people
		totalAmountOfPeople += numberOfPeople
	}

	// Find percentages for each group
	group1 := float64(countMusala) / float64(totalAmountOfPeople) * 100
	group2 := float64(countMontblanc) / float64(totalAmountOfPeople) * 100
	group3 := float64(countKilimandjaro) / float64(totalAmountOfPeople) * 100
	group4 := float64(countK2) / float64(totalAmountOfPeople) * 100
	group5 := float64(countEverest) / float64(totalAmountOfPeople) * 100

	// Print output, don't forget new line and %-tage
	fmt.Printf("%.2f%%\n", group1)
	fmt.Printf("%.2f%%\n", group2)
	fmt.Printf("%.2f%%\n", group3)
	fmt.Printf("%.2f%%\n", group4)
	fmt.Printf("%.2f%%\n", group5)
}
