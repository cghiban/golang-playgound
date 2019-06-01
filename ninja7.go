package main

import (
	"fmt"
)

type person struct {
	name, address string
}

func changeAddress(p *person) {
	// both ways works: x.f is shorthand for (*x).f
	//p.address = "secret address"
	(*p).address = "secret address"
}

func main() {
	p := person{
		name:    "A B",
		address: "home",
	}

	fmt.Printf("%+v\n", p)
	changeAddress(&p)
	fmt.Printf("%+v\n", p)
}
