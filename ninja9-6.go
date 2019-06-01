package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpus: ", runtime.NumCPU())
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("cparchus: ", runtime.GOARCH)
}
