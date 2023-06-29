package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	workingWithMaps()

}

func workingWithSlices() {

	var s []int
	s = []int{1, 2, 3}
	s = append(s, 4, 5)
	s = slices.Delete(s, 0, 2)
	fmt.Println(s)
}

func workingWithMaps() {
	var m map[string]int
	m = map[string]int{"foo": 1, "bar": 2}

	if v, ok := m["foo"]; ok {
		fmt.Println("found", v)
	}

}

func menuItems() {
	var m map[string][]string
	m = map[string][]string{
		"coffee": {"Coffee", "Espresso"},
		"tea":    {"Hot Tea", "chai"},
	}
	fmt.Println(m)
}

func workingWithStructs() {
	var s struct {
		name string
		id   int
	}
	s.name = "John"
	s.id = 1
	fmt.Println(s)
}

func workingWithTypedStruct() {
	type myStruct struct {
		name string
		id   int
	}

	var s myStruct
	var y myStruct

	s = myStruct{
		name: "Jesus",
		id:   43,
	}

	y = myStruct{
		name: "yeo",
		id:   54,
	}
	fmt.Println(s, y)
}

func personStruct() {

	type Person struct {
		name string
		age  int
	}

	person1 := Person{"John", 12}
	x := Person{name: "James", age: 42}
	fmt.Println(person1, x)

}
