package main

import (
	"fmt"
)

func main() {

	workingWithConst()
	workingWithIota()
	workingWithPointers()
}

func workingWithConst() {
	const a = 43
	const b string = "hello, world"
	const (
		d = true
		c = "foo"
	)
	fmt.Println(c)
}

func workingWithIota() {
	const a = iota
	const (
		b = iota
		c
		d = 3 * iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func workingWithPointers() {
	a := 43
	b := &a
	*b = 20
	fmt.Println(a)

	b = new(int)

	fmt.Println(b, *b, a)
}
