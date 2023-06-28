package main

import "fmt"

func main() {

	t, g, c := 3, 4, 5
	changeValueAtPointer(&t, &g, &c)

	fmt.Println(t, g, c)
}

func changeValueAtPointer(a ...*int) {

	for _, v := range a {
		*v = 10
	}
}
