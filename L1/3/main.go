// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.
package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	Pow(a)
}

func Pow(arr []int) {
	b := make(chan int, len(arr))
	for i, _ := range arr {
		go func(b chan int, a int) {
			b <- a * a
		}(b, arr[i])
	}
	var res int
	for i := 0; i < len(arr); i++ {
		res += <-b
	}
	fmt.Println(res)
}
