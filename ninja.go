package main

import (
	"fmt"
	"runtime"
)

var x int = 42
var y string = "James B"
var z bool = true

type beta int

var b beta = 42
var c int

func e2() {
	fmt.Println(x, y, z)

	s := fmt.Sprintf("x=%v\ty=%q\tz=%v", x, y, z)
	fmt.Println(s)
}

func e4() {
	fmt.Println("b =", b)
}

func e5() {
	fmt.Printf("b = %v; type =%T\n", b, b)
	c = int(b)
	fmt.Println("c =", c)
}

func main() {
	e2()
	e4()
	e5()

	fmt.Println(runtime.GOARCH)
	s := "Ana are mere pădurețe"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}
	for i, v := range s {
		fmt.Printf("%d = %#x\n", i, v)
	}
}
