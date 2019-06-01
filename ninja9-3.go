package main

import (
	"fmt"
	"runtime"
	"sync"
)

var g int
var z int //
var wg sync.WaitGroup

func incvar() {
	x := g
	runtime.Gosched() // yields the cpu to do someth else
	x++
	g = x
	wg.Done()
}

func main() {
	fmt.Println("cpus: ", runtime.NumCPU())
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	numnum := 100
	wg.Add(numnum)

	for i := 0; i < numnum; i++ {
		go incvar()
	}

	wg.Wait()
	fmt.Println("g =", g)

	fmt.Println("Done!")
}
