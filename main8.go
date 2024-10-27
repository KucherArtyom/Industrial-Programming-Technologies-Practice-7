package main

import (
	"fmt"
)

func findMax(arr []int) int {
	if len(arr) == 0 {
		fmt.Println("Массив пустой")
		return 0
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	var arrayAmount int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&arrayAmount)

	array := make([]int, arrayAmount)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < arrayAmount; i++ {
		fmt.Scan(&array[i])
	}

	maxNumber := findMax(array)
	fmt.Println("Максимальный элемент в массиве:", maxNumber)
}
