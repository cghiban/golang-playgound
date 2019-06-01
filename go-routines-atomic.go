package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("cpus: ", runtime.NumCPU())
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < gs; i++ {
		go func() { // funcliteral
			atomic.AddInt64(&counter, 1)
			fmt.Println("Count: ", atomic.LoadInt64(&counter))
			//runtime.Gosched() // yield = go ahead and do somt else
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	fmt.Println("counter: ", counter)
}
