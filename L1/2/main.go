// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.
package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	Pow(a)
}

func Pow(arr []int) {
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for i, _ := range arr {
		go func(a *int) {
			defer wg.Done()
			*a *= *a
		}(&arr[i])
	}
	wg.Wait()
	fmt.Fprint(os.Stdout, arr)
}
