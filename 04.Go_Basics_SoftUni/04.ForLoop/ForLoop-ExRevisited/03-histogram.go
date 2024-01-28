package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	// Параметри р1 под 200, р2 до 399, р3 до 599, р4 до 799, р5 от 800 нагоре
	// Търсим %-тите на целите числа попадащи във всяка категория
	// Следователно търсим бройката във всяка една категория в съотношение с всички числа

	var p1, p2, p3, p4, p5 int

	for i := 0; i < n; i++ {
		var number int
		fmt.Scanln(&number)

		// Намираме кое число в коя категория се намира
		if number < 200 {
			p1++
		} else if number < 400 {
			p2++
		} else if number < 600 {
			p3++
		} else if number < 800 {
			p4++
		} else if number >= 800 {
			p5++
		}
	}

	// След приключване на цикъла трябва да намерим процентите
	percentP1 := float64(p1) / float64(n) * 100
	percentP2 := float64(p2) / float64(n) * 100
	percentP3 := float64(p3) / float64(n) * 100
	percentP4 := float64(p4) / float64(n) * 100
	percentP5 := float64(p5) / float64(n) * 100

	// Print output
	fmt.Printf("%.2f%%\n", percentP1) // Има нужда от двоен % за да принтираме %-та като символ
	fmt.Printf("%.2f%%\n", percentP2)
	fmt.Printf("%.2f%%\n", percentP3)
	fmt.Printf("%.2f%%\n", percentP4)
	fmt.Printf("%.2f%%\n", percentP5)
}
