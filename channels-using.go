package main

import "fmt"

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	ch := make(chan int)

	go foo(ch) // send
	bar(ch)    // receive

	fmt.Println("About to exit")
}
