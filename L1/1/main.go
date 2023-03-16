// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	human Human
}

func (a *Human) SetName(name string) {
	a.name = name
	fmt.Printf("Your name is %s", name)
}

func main() {
	a := &Action{}
	a.human.SetName("Vasiliy")
}
