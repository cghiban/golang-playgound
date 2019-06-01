package main

import "fmt"

func main() {

	jb := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   66,
	}

	fmt.Printf("%v - %T\n", jb, jb)
}
