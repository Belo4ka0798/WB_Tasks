// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World! Привет, Мир!"
	arr := strings.Split(str, " ")

	for i := 0; i <= len(arr)-1; i++ {
		tmp := []rune(arr[i])
		for i := len(tmp) - 1; i >= 0; i-- {
			fmt.Print(string(tmp[i]))
		}
		fmt.Print(" ")
	}
}
