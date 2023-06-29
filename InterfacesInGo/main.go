package main

import "fmt"

func main() {

}

type Animal interface {
	Speak()
}

type Cat struct {
	numLegs int
}

type Dog struct {
	numLegs int
}

type Kangaroo struct {
	numPouches int
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}
func (c Dog) Speak() {
	fmt.Println("Woof")
}
func (c Kangaroo) Speak() {
	fmt.Println("Boinggggg")
}

func typeAssertions() {
	var c, d, e int
	var f, g, h string
}
