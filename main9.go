package main

import (
	"fmt"
)

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var firstNumber, secondNumber int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&secondNumber)

	result := gcd(firstNumber, secondNumber)
	fmt.Printf("Наибольший общий делитель чисел %d и %d: %d\n", firstNumber, secondNumber, result)
}
