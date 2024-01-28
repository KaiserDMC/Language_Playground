package main

import "fmt"

func main() {
	// Проверка за вход - няма входни данни

	// Параметри:
	// Диапазон от 1 до 1000, цели числа
	// Цикъл с 3 параметъра: начало, проверка, промяна/стъпка

	for number := 1; number <= 1000; number++ {
		if number%10 == 7 {
			fmt.Println(number)
		}
	}

	// Оптимизация - slower in Judge

	for i := 7; i <= 997; i += 10 {
		fmt.Println(i)
	}

}
