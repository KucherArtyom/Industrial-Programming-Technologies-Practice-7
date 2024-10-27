package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int
	fmt.Print("Введите число для проверки на простоту: ")
	fmt.Scan(&number)

	if isPrime(number) {
		fmt.Println("Число является простым.")
	} else {
		fmt.Println("Число не является простым.")
	}
}
