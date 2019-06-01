package main

import (
	"fmt"
)

func main() {

	const (
		a         = 42
		b float64 = 3.2
		c         = iota
		d
	)

	const (
		e = iota
		f
	)

	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println("c =", c)
	fmt.Printf("%T\n", c)

	fmt.Println("d =", d)
	fmt.Printf("%T\n", d)

	fmt.Println("e =", e)
	fmt.Printf("%T\n", e)

	y := f << 2
	fmt.Printf("f = %d\t\t%8b\n", f, f)
	fmt.Printf("y = %d\t\t%8b\n", y, y)
	//fmt.Printf("%T\n", f)
}
