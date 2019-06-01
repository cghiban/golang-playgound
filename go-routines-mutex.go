package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doSomething(x int) int {
	return x * 5
}

func main() {
	fmt.Println("cpus: ", runtime.NumCPU())
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(100)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() { // funcliteral
			mu.Lock()
			v := counter
			//time.Sleep(1)
			runtime.Gosched() // yield = go ahead and do somt else
			v++
			counter = v
			mu.Unlock()
			//fmt.Println("goroutines: ", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	fmt.Println("counter: ", counter)
	return
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}
