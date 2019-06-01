package main

import "fmt"

func z1(a int) {
	a = 33
}

func z2(a *int) {
	*a = 33
}

func main() {
	a := 22
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(&a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)  // b is an address
	fmt.Println(*b) // what's at the address b
	fmt.Println(*&a)

	*b = 11 // *b ==  == 11

	x := 0
	z1(x)
	fmt.Println("x after z1(int):", x)

	z2(&x)
	fmt.Println("x after foo(*int):", x)
}
