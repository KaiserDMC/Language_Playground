package main

import (
	"fmt"
	"math"
)

func main() {
	// Read input, Lily's years n
	var n int
	fmt.Scanln(&n)

	// Odd years she gets a toy, in the end she sells the toys for P lev
	// Even years she gets money - 10 lev per year, her brother takes 1 lev from her each year
	// Check if she can afford a new washing machine

	var washerPrice, sumOfMoney float64
	var priceOfEachToy, moneyPerYear int

	// Read the rest of the input
	fmt.Scanln(&washerPrice)
	fmt.Scanln(&priceOfEachToy)

	for i := 1; i <= n; i++ {
		// If even years, she gets money
		if i%2 == 0 {
			moneyPerYear += 10
			sumOfMoney += float64(moneyPerYear)
			sumOfMoney -= 1
		} else {
			// If odd she gets a toy, but sells it in the end
			sumOfMoney += float64(priceOfEachToy)
		}
	}

	// Check if she has enough money
	differenceInPrice := sumOfMoney - washerPrice

	if differenceInPrice >= 0 {
		// She has enough money to buy the washing machine
		fmt.Printf("Yes! %.2f", differenceInPrice)
	} else {
		// She does not have enough
		fmt.Printf("No! %.2f", math.Abs(differenceInPrice))
	}
}
