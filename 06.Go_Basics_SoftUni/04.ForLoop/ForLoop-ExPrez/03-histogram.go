package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var p1, p2, p3, p4, p5 int = 0, 0, 0, 0, 0
	for i := 0; i < n; i++ {
		var number int
		fmt.Scanln(&number)

		if number < 200 {
			p1++ // p1+=1
		} else if number < 400 {
			p2++
		} else if number < 600 {
			p3++
		} else if number < 800 {
			p4++
		} else {
			p5++
		}
	}

	percentP1 := float64(p1) / float64(n) * 100
	percentP2 := float64(p2) / float64(n) * 100
	percentP3 := float64(p3) / float64(n) * 100
	percentP4 := float64(p4) / float64(n) * 100
	percentP5 := float64(p5) / float64(n) * 100

	fmt.Printf("%.2f%%\n", percentP1)
	fmt.Printf("%.2f%%\n", percentP2)
	fmt.Printf("%.2f%%\n", percentP3)
	fmt.Printf("%.2f%%\n", percentP4)
	fmt.Printf("%.2f%%\n", percentP5)

}
