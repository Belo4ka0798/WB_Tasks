// Разработать программу,
// которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan int)
	CheckType(21)
	CheckType(3.14)
	CheckType("Hello, World!")
	CheckType(true)
	CheckType(ch)
}

func CheckType(a any) {
	ty := reflect.TypeOf(a).Kind()
	switch ty {
	case reflect.Int:
		fmt.Println("Its INTEGER")
		break
	case reflect.Float32:
		fmt.Println("Its FLOAT32")
	case reflect.Float64:
		fmt.Println("Its FLOAT64")
		break
	case reflect.String:
		fmt.Println("Its STRING")
		break
	case reflect.Bool:
		fmt.Println("Its BOOL")
		break
	case reflect.Chan:
		fmt.Println("Its CHANNELS")
		break
	default:
		fmt.Println("I dont know this type!")
		break
	}
}
