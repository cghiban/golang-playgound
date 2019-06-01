package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	//
}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("num CPUs:", runtime.NumCPU())
	fmt.Println("num goroutines:", runtime.NumGoroutine())

	wg.Add(1) // one thing to wait for

	go foo()
	bar()

	fmt.Println("num goroutines:", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo:", i)
	}

	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar:", i)
	}
	//wg.Done()
}
