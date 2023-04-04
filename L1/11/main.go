// Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr1, arr2 [20]int
	m := make(map[int]int)
	rand.Seed(time.Now().UnixNano())
	arr1[0] = 1
	arr2[0], arr2[1] = 1, 1
	for i := 0; i < len(arr1); i++ {
		m[arr1[i]] = 1
	}
	for i := 0; i < len(arr2); i++ {
		if _, ok := m[arr2[i]]; ok {
			m[arr2[i]]++
		}
		if m[arr2[i]] == 2 {
			fmt.Println(arr2[i])
		}
	}
}
