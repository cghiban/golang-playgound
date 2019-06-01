package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func doSomething(x int) int {
	return x * 5
}

//
func main() {
	fmt.Println("cpus: ", runtime.NumCPU())
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < gs; i++ {
		go func() { // funcliteral
			v := counter
			time.Sleep(1)
			runtime.Gosched() // yields the cpu to do someth else
			v++
			counter = v
			//fmt.Println("goroutines: ", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	fmt.Println("counter: ", counter)
}
