// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64 = 6782
	bit := 5
	fmt.Printf("Result: %d\n", ChangeBits(num, bit))

}

func ChangeBits(num int64, i int) int64 {
	var tmp []byte = []byte(strconv.FormatInt(int64(num), 2))
	fmt.Printf("Begin : %s\n", string(tmp))
	for iter, _ := range tmp {
		if iter == i && tmp[i-1] == 49 {
			tmp[i-1] = 48
		} else if iter == i && tmp[i-1] == 48 {
			tmp[i-1] = 49
		}
	}
	res, _ := strconv.Atoi(string(tmp))
	return int64(res)
}
