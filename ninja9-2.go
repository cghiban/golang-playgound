package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Printf("Person %s is speaking..", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{name: "Dragnea"}
	//saySomething(p) // is is not going to work
	saySomething(&p)
}
