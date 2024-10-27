package main

import (
	"fmt"
)

func sumOfSquares(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i * i
	}
	return sum
}

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	result := sumOfSquares(number)
	fmt.Printf("Сумма квадратов чётных чисел от 1 до %d: %d\n", number, result)
}
