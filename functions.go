package main

import "fmt"

// variadic parameters
func foo(x ...int) int {
	// x is a slice of int, []int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

func boo() {
	fmt.Println(`boo -- defered`)
}

func far() {
	fmt.Println("far")
}

// this func returns a func that returns an int
func returnFunction() func() int {
	return func() int {
		return 5
	}
}

// this func returns a func that returns an int
func adder(x int) func(int) int {
	return func(z int) int {
		return x + z
	}
}

func main() {
	fmt.Println(foo())
	fmt.Println(foo(1, 3, 4))

	// slice of int
	xi := []int{1, 2, 3, 4}
	xiSum := foo(xi...) // unfurling :)
	fmt.Println(xiSum)

	fmt.Println("//------ DEFER --------------")
	// defer
	defer boo()
	far()

	fmt.Println("//------ ANONYMOUS --------------")
	//anonimous functions:
	func() {
		fmt.Println("Anonymous function called :D")
	}()
	func(x, y int) {
		fmt.Println("Anonymous function called :D")
		fmt.Println("s =", y+x)
	}(3, 4)

	f := func() {
		fmt.Println("my first func expression")
	}
	fmt.Printf("f is a «%T»\n", f)
	f()

	five := returnFunction()
	fmt.Println("five:", five)
	fmt.Println("five():", five())

	add3 := adder(3)
	add99 := adder(99)
	fmt.Println("add3(5):", add3(5))
	fmt.Println("add99(1):", add99(1))

	fmt.Println("adder(9)(11):", adder(9)(11))

}
