// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.
package main

import "fmt"

func main() {
	str := "Hello, World! Привет, Мир!"
	tmp := []rune(str)
	for i := len(tmp) - 1; i >= 0; i-- {
		fmt.Print(string(tmp[i]))
	}
}
