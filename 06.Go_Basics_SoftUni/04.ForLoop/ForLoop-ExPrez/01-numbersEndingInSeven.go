package main

import "fmt"

func main() {
	// Read Input - no input

	for i := 1; i <= 1000; i++ {
		if i%10 == 7 {
			fmt.Println(i)
		}
	}

	for j := 7; j <= 997; j++ {
		if j%10 == 7 {
			fmt.Println(j)
		}
	}

	for k := 7; k <= 997; k += 10 {
		fmt.Println(k)
	}
}
