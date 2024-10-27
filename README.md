# Industrial-Programming-Technologies-Practice-7 (ЭФМО-02-24 Кучер Артем Сергеевич)
## Список задач на тему Функции. Язык программирования go.
### 1.Функция вычисления площади треугольника
```
package main

import (
	"fmt"
)

func triangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}

func main() {
	var base, height float64
	fmt.Print("Введите длину основания треугольника: ")
	fmt.Scan(&base)
	fmt.Print("Введите высоту треугольника: ")
	fmt.Scan(&height)
	area := triangleArea(base, height)
	fmt.Printf("Площадь треугольника: %v\n", area)
}
```
### 2.Сортировка массива по возрастанию
```
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
```
### 3.Сумма квадратов чётных чисел
```
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
```
### 4.Проверка палиндрома
```
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var filteredRunes []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			filteredRunes = append(filteredRunes, r)
		}
	}

	n := len(filteredRunes)
	for i := 0; i < n/2; i++ {
		if filteredRunes[i] != filteredRunes[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var line string
	fmt.Print("Введите строку для проверки на палиндром: ")
	fmt.Scanln(&line)

	if isPalindrome(line) {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}
```
### 5.Проверка числа на простоту
```
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
```
### 6.Генерация последовательности простых чисел
```
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

func generatePrimes(limit int) []int {
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var limit int
	fmt.Print("Введите предел для генерации простых чисел: ")
	fmt.Scan(&limit)

	primes := generatePrimes(limit)
	fmt.Println("Простые числа до предела:", primes)
}
```
### 7.Перевод числа в двоичную систему
```
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
```
### 8.Нахождение максимального элемента в массиве
```
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

	numbers := make([]int, arrayAmount)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < arrayAmount; i++ {
		fmt.Scan(&numbers[i])
	}

	maxNumber := findMax(numbers)
	fmt.Println("Максимальный элемент в массиве:", maxNumber)
}
```
### 9.Функция вычисления НОД (наибольший общий делитель)
```
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
```
### 10.Сумма элементов массива
```
package main

import (
	"fmt"
)

func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	var arrayNumber int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&arrayNumber)

	numbers := make([]int, arrayNumber)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < arrayNumber; i++ {
		fmt.Scan(&numbers[i])
	}

	sum := sumArray(numbers)
	fmt.Println("Сумма элементов массива:", sum)
}
```
