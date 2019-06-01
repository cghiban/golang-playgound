package main

import "fmt"

//send
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//q <- 0
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from even: ", v)
		case v, ok := <-o:
			fmt.Println("from odd: ", v, ok)
		case v, ok := <-q:
			if !ok {
				fmt.Println("from quit: ", v, ok)
				return
			} else {
				fmt.Println("from quit: ", v, ok)
			}
		}
	}
}

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("Done")
}
