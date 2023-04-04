// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

func main() {
	ex := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve"}
	m := make(map[string]int, 11)
	wg := sync.WaitGroup{}
	mut := sync.Mutex{}
	for i, _ := range ex {
		wg.Add(1)
		go func(a map[string]int, d string, c int, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			a[d] = c
			mu.Unlock()
		}(m, ex[i], i, &mut)
	}
	wg.Wait()
	fmt.Println(m)
}
