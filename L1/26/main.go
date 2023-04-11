// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
package main

import (
	"fmt"
	"unicode"
)

func isUnique(s *string) bool {
	set := make(map[rune]uint8)
	for _, value := range *s {
		if unicode.IsLetter(value) {
			value = unicode.ToLower(value)
		}
		_, contains := set[value]
		if contains {
			return false
		}
		set[value] = 0
	}
	return true
}

func main() {
	s1 := "Hi man😎"
	s2 := "🌚Hello world"
	s3 := "abcd"
	s4 := "abCdefAaf"
	s5 := "aabcd"

	fmt.Println(isUnique(&s1))
	fmt.Println(isUnique(&s2))
	fmt.Println(isUnique(&s3))
	fmt.Println(isUnique(&s4))
	fmt.Println(isUnique(&s5))
}
