package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	if res, err := throwAnErrorIfOdd(2); err == nil {
		fmt.Println(res)

	}

	if res, err := throwAnErrorIfOdd(1); err == nil {
		fmt.Println(res)
	}

	res, err := divdide(10, 0)
	fmt.Println(res, err)

	divdide2(10, 0)
}

func readFromFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Couldn't open file")
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func throwAnErrorIfOdd(num int) (int, error) {

	if num%2 != 0 {
		e := errors.New("This is an error")
		return num, e
	}
	return num, nil

}
func divdide(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("Invalid Divisior")
	}
	return l / r, nil
}
func divdide1(l, r int) int {

	return l / r
}
func divdide2(l, r int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l / r, nil
}
