package main

import "fmt"

func main() {

	const (
		y1 = iota + 2019
		y2
		y3
	)

	x := 21
	y := x << 1
	z := x << 2
	fmt.Printf("%d\t%8b\t%#X\n", x, x, x)
	fmt.Printf("%d\t%8b\t%#X\n", y, y, y)
	fmt.Printf("%d\t%8b\t%#X\n", z, z, z)

	a := `some 
	raw string
	literal`
	fmt.Println(a)

	fmt.Printf("%d\t%d\t%d\n", y1, y2, y3)
}
