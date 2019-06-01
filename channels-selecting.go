package main

import "fmt"

//send
func send(e chan int, o chan int, q chan<- int) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e chan int, o chan int, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from even: ", v)
		case v := <-o:
			fmt.Println("from odd: ", v)
		case v := <-q:
			fmt.Println("from quit: ", v)
			//close(q)
			return
		}
	}
}

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("Done")
}
