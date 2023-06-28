package main

import "fmt"

func main() {
	var s []int
	s = []int{1, 2, 3}
	loopThroughArray(s)
	var m map[string]string
	m = map[string]string{"foo": "inFoo", "bar": "inBar"}
	loopThroughMap(m)
}

func workingWithInfiniteLoop() {

	for {
		fmt.Println("In an Infinite Loop")
	}
}

func loopingTillCondition(num int) {

	for num > 0 {
		fmt.Println(num)
		num--
	}
}

func workingWithInitalizedLoop(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func loopThroughArray(arr []int) {
	for k, v := range arr {
		fmt.Println(k, v)
	}
}
func loopThroughMap(myMap map[string]string) {
	for _, v := range myMap {
		fmt.Println(v)
	}
}
