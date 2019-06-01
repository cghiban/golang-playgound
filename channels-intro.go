package main

import "fmt"

func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	// ch := make(chan int, 1) // buffered channel
	go func() {
		ch <- doSomething(5)
		fmt.Println("done with the channel")
	}()
	fmt.Println(<-ch)
}
