package main

import (
	"fmt"
	"sync"
)

func main() {
	workingWithChannelLoops()
}
func workingWithChannelLoops() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
func workingWithChannelOperation() {
	ch1, ch2 := make(chan string), make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("No Messages Recieved")
	}
}
func sendingMessagesWithChannel() {
	ch := make(chan string)
	var wg2 sync.WaitGroup

	wg2.Add(1)
	go func() {
		ch <- "In the first function"
	}()
	go func() {
		fmt.Println(<-ch)
		wg2.Done()
	}()
	wg2.Wait()
}

func declaringAChannel() {
	ch := make(chan string)

	ch <- "Sending a message" //sending a message

	msg := <-ch
	fmt.Println(msg)

}
func multipleWaitGroups() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("This is async")

	}()

	go func() {
		fmt.Println("This also is async")
		wg.Done()
	}()

	fmt.Println("This is sync")
	wg.Wait()
}

func declaringAGoRoutine() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("do something")
		wg.Done()
	}()
	wg.Wait()
}
