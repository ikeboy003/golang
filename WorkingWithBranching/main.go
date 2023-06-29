package main

import "fmt"

func main() {
	workingWithGoToStatements(4)
	workingWithGoToStatements(6)

}

func workingWithIfStatement() {

}

func workingWithSwitch() {

	u := 5
	switch u {
	case 1:
		fmt.Println("1")
	case 2, 5:
		fmt.Println("2 or 5")
	}
}

func workingWithLogicalSwith() {
	u := 5
	j := 3

	switch true {
	case u > 7 || j < 4:
		fmt.Println("greater")
	case u < 4:
		fmt.Println("less")
	}
}

func workingWithDeferredFunction(num int) {
	defer fmt.Println("Deferred Function ", num)
	workingWithLogicalSwith()

}

func divide(dividend, divisor int) int {

	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}

	}()
	return dividend / divisor
}

func workingWithGoToStatements(numToExit int) {

	if numToExit > 5 {
		goto exit
	}
	fmt.Println("this should run if it is less than 5")
exit:
	fmt.Println("We are now exiting the program")

}
