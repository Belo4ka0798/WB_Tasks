// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	//TODO Сделать 22 задание
	a := new(big.Int)
	b := new(big.Int)

	aStr := ""
	bStr := ""
	fmt.Printf("Enter first number: ")
	fmt.Scan(&aStr)
	if _, ok := a.SetString(aStr, 10); ok {
		fmt.Println("Wrong data")
		os.Exit(1)
	}
	fmt.Printf("Enter second number: ")
	fmt.Scan(&bStr)
	if _, ok := b.SetString(bStr, 10); ok {
		fmt.Println("Wrong data!")
	}
	result := new(big.Int)
	status := false
	for !status {
		operation := ""
		fmt.Printf("Select an operation (+, -, *, /): ")
		fmt.Scan(&operation)
		status = true
		switch operation {
		case "+":
			result.Add(a, b)
		case "-":
			result.Sub(a, b)
		case "*":
			result.Mul(a, b)
		case "/":
			zero := big.NewInt(0)
			if b.Cmp(zero) == 0 {
				fmt.Println("Divison by zero, select another operation")
				status = false
			} else {
				result.Div(a, b)
			}
		default:
			fmt.Println("Wrong operation!")
			status = false
		}
	}
	fmt.Println(result)
}
