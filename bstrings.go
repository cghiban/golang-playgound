package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOARCH)
	s := "Ana are mere pădurețe"
	bs := []byte(s)
	for i := 0; i < len(s); i++ {
		// UTF-8
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("\n", bs)
	fmt.Printf("%T\n", bs)

	for i, v := range s {
		fmt.Printf("at index %d we have hex %#x or %v\n", i, v, bs[i])
	}
}
