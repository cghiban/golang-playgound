package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("cpus: ", runtime.NumCPU())
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	numPrints := 5
	var wg sync.WaitGroup
	wg.Add(numPrints)

	for i := 0; i < numPrints; i++ {
		go func(arg int) { // funcliteral
			fmt.Println("i: ", arg)
			runtime.Gosched() // yield = go ahead and do somt else
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("goroutines: ", runtime.NumGoroutine())
	fmt.Println("Done!")
}
