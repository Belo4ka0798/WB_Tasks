// Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"fmt"
	"sort"
)

func main() {
	sea := 0
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 4, 5, 1, 3, 6, 8, 4, 123, 6, 786, 66, 23123, 67786, 8, 9, 5445, 2, 23, 2, 67, 98, 89, 55}
	check := sort.Search(len(arr), func(ind int) bool { return arr[ind] >= sea })
	if check < len(arr) && arr[check] == sea {
		fmt.Println("Number is Find!")
		return
	} else {
		fmt.Println("Number is not Find!")
	}
}
