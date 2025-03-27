package main

import (
	"fmt"
)

func main() {
	// Read Input, number of tournaments n, initial points, for each tournament-> placement
	var numTournaments, initialPoints int
	fmt.Scanln(&numTournaments)
	fmt.Scanln(&initialPoints)

	var points, numWins int // Separate variable for his tournament points
	// For each tournament
	for i := 0; i < numTournaments; i++ {
		// Get his placement
		var placement string
		fmt.Scanln(&placement)

		// Give point based on his placement
		switch placement {
		case "W":
			points += 2000
			numWins++
		case "F":
			points += 1200
		case "SF":
			points += 720
		}
	}
	// Get his total amount of points, initial and after all tournaments
	totalPoints := initialPoints + points
	// Get average points from tournaments, either int or floor
	averageTournamentPoints := points / numTournaments
	//averageTournamentPointsFloor := math.Floor(float64(points) / float64(numTournaments))
	percentageWins := float64(numWins) / float64(numTournaments) * 100

	// Print all required outputs
	fmt.Printf("Final points: %v\n", totalPoints)
	fmt.Printf("Average points: %v\n", averageTournamentPoints)
	//fmt.Printf("Average points: %.0f\n", averageTournamentPointsFloor)
	fmt.Printf("%.2f%%", percentageWins)
}
