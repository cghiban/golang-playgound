package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	eu := person{
		first: "Him",
		last:  "Himself",
		age:   30,
	}

	miaua := person{
		first: "Her",
		last:  "Herself",
	}

	fmt.Println(eu, miaua)
	fmt.Println(eu.first, miaua.first)
}
