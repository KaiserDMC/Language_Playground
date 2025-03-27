package main

import (
	"fmt"
	"math"
)

func main() {
	// Input
	var n int
	fmt.Scanln(&n)

	maxNumber := math.MinInt64 // math.MinInt
	var sum int

	for i := 0; i < n; i++ {
		var currentNumber int
		fmt.Scanln(&currentNumber)

		if currentNumber > maxNumber {
			maxNumber = currentNumber
		}

		sum += currentNumber
	}

	sumWithoutMaxNumber := sum - maxNumber

	if sumWithoutMaxNumber == maxNumber {
		fmt.Println("Yes")
		fmt.Printf("Sum = %v", maxNumber)
	} else {
		fmt.Println("No")
		fmt.Printf("Diff = %.0f", math.Abs(float64(sumWithoutMaxNumber)-float64(maxNumber)))
	}
}
