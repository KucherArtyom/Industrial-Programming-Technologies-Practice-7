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
