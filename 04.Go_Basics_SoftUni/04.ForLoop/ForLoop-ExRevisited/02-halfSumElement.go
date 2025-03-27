package main

import (
	"fmt"
	"math"
)

func main() {
	// Прочитаме входните данни
	// Число н последвани от н на брой числа
	var n int
	fmt.Scanln(&n)

	maxNumber := math.MinInt64
	var sum int

	for i := 0; i < n; i++ {
		// Прочинаме всяко число
		var currentNumber int
		fmt.Scanln(&currentNumber)

		// Проверка дали сегашното число е по-голямо от сегашното най-голямо число
		if currentNumber > maxNumber {
			maxNumber = currentNumber
		}

		// Намираме сумата при всеки цикъл
		sum += currentNumber
	}

	// След приключване на цикъла
	sumWithoutMax := sum - maxNumber

	if sumWithoutMax == maxNumber {
		// Ако сумата без най-голямото е равна на най-голямото
		fmt.Println("Yes")
		fmt.Printf("Sum = %v", sumWithoutMax)
	} else {
		// Ако не са равни
		fmt.Println("No")
		fmt.Printf("Diff = %.0f", math.Abs(float64(sumWithoutMax)-float64(maxNumber))) // Модулно, за да няма отрицателна разлика
		// Проблем при %v в Judge
	}
}
