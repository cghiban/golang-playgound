package main

import (
	"fmt"
	"runtime"
	"sync"
)

var z int //
var wg sync.WaitGroup
var mu sync.Mutex

func incvarWMutex() {
	mu.Lock()
	x := z
	// no need for runtime.Gosched()
	//runtime.Gosched()
	x++
	z = x
	mu.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("cpus: ", runtime.NumCPU())
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	numnum := 100
	// now w/a mutex
	wg.Add(numnum)

	for i := 0; i < numnum; i++ {
		go incvarWMutex()
	}
	wg.Wait()
	fmt.Println("z =", z)

	fmt.Println("Done!")
}
