package main

import "fmt"

func main() {
	s := []string{
		"Hello",
		"My Name is",
	}
	t := []int{1, 3, 4}

	genericPrint(s)
	genericPrint(t)

	str, num := "Hello", 3
	newNum, newStr := genericSwap(str, num)
	fmt.Println(newNum, newStr)

}

func genericPrint[T any](s []T) {

	for _, v := range s {
		fmt.Println(v)
	}

}

func genericSwap[T, V any](f T, g V) (V, T) {

	return g, f
}
