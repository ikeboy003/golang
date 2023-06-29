package main

import "fmt"

func main() {

	a1 := []int{1, 2, 3}

	a2 := []float64{1.2323, 2.234, 3.342}
	a3 := []string{"Hi", "My name is", "Slim Shady"}
	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)
	fmt.Println(s1, s2, s3)
	cusStr := customString("HMMMM")
	a4 := []customString{customString("Hello"), customString("How are you"), customString("My name is"), cusStr}
	s4 := add(a4)
	fmt.Println(s4)

}

type addable interface {
	~string | int | float64
}

func add[T addable](vals []T) T {
	var sum T
	for _, v := range vals {
		sum += v
	}
	return sum
}

type customString string // aliasing a type
