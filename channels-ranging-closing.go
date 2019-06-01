package main

import "fmt"

//send
func foo(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- 42 + i
	}
	close(c)
}

func main() {
	ch := make(chan int)

	go foo(ch) // send

	// receive
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Done")

}
