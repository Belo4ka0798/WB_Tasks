// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 4, 5, 1, 3, 6, 8, 4, 123, 6, 786, 66, 23123, 67786, 8, 9, 5445, 2, 23, 2, 67, 98, 89, 55}
	sort.Ints(arr)
	fmt.Println(arr)
}
