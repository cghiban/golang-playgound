package main

import "fmt"

func doSomething(x int) int {
	return x * 5
}

func main() {
	//ch := make(chan int)
	ch := make(chan int, 2) // buffered channel, send and receive channel
	cs := make(chan<- int)  // send only channel
	cr := make(<-chan int)  // receive nly channel

	ch <- 42
	ch <- 43

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("-----")
	fmt.Printf("ch type: %T\n", ch)
	fmt.Printf("cs type: %T\n", cs)
	fmt.Printf("cr type: %T\n", cr)
}
