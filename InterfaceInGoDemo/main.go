package main

import "fmt"

func main() {
	workingWithSwitchTypeAssertions()
}

type animal interface {
	Speak()
}

type dog struct {
	numLegs int
}

type cat struct {
	numLegs int
}

func (c cat) Speak() {
	fmt.Println("Meoooooow")
}

func (c dog) Speak() {
	fmt.Println("Woof")
}

func workingWithIfTypeAssertion() {
	c := cat{3}
	var a animal = c

	var d1 dog
	_, ok := a.(dog)

	if ok {
		d1 = a.(dog)
		fmt.Println("its a dog")
	} else {
		fmt.Println("Not a dog")
		fmt.Println(d1)
	}

}

func workingWithSwitchTypeAssertions() {
	c := cat{3}
	var a animal = c

	switch a.(type) {
	case dog:
		fmt.Println("its a dog")
	case cat:
		fmt.Println("its a cat")
	}

}
