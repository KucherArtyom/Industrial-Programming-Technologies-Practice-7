package main

import (
	"fmt"
)

func toBinary(n int) string {
	if n == 0 {
		return "0"
	}
	binary := ""
	for n > 0 {
		bit := n % 2
		binary = fmt.Sprintf("%d%s", bit, binary)
		n /= 2
	}
	return binary
}

func main() {
	var number int
	fmt.Print("Введите число для перевода в двоичную систему: ")
	fmt.Scan(&number)

	numberBinary := toBinary(number)
	fmt.Printf("Число %d в двоичной системе: %s\n", number, numberBinary)
}
