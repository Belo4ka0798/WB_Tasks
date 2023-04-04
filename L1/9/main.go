// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.
package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	ch1 := make(chan int, len(arr)-1)
	ch2 := make(chan int, len(arr)-1)

	for _, d := range arr {
		ch1 <- d
		go func() {
			ch2 <- 2 * <-ch1
		}()
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(<-ch2)
	}
}
