package main

import "fmt"

func main() {
	str, str2 := "Ikenna", "nnamdi"
	funcWithPointer(&str)
	funcWithoutPointer(str2)

	fmt.Println(str, str2)

	x, y, z := funcWithMultipleReturns()
	fmt.Println(x, y, z)

	funcWithVaradicParams("Hi", "How Are You", "Doing Today")

}

func workingWithFunctions() {
	fmt.Println("This is a Function")
}

func funcWithParams(param, param2 int, param3, param4 string) string {
	fmt.Println(param2, param, param3, param4)
	return ""
}
func functionWithReturn() string {
	return "Hello"
}

func funcWithPointer(name *string) {

	*name = "New Name"
}
func funcWithoutPointer(name string) {
	name = "hello"
}
func funcWithVaradicParams(strings ...string) {
	for i, v := range strings {
		fmt.Println(i, v)
	}
}
func funcWithMultipleReturns() (mystr string, mystr1 string, mystr2 string) {
	return
}
