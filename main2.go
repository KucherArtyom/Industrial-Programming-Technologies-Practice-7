package main

import (
	"fmt"
)

func sortArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	var arrayAmount int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&arrayAmount)

	numbers := make([]int, arrayAmount)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < arrayAmount; i++ {
		fmt.Scan(&numbers[i])
	}
	sortedArr := sortArray(numbers)
	fmt.Println("Отсортированный массив:", sortedArr)
}
