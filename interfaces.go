package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	country string
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- a person")
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "- a secret agent")
}

// any type that has a speak method is a human type also
// a vaue can be of more than one type
type human interface {
	speak()
}

// will use a switch on type
func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("person:", h.(person).first) // using assertion
	case secretAgent:
		fmt.Println("secretAgent:", h.(secretAgent).first)
	default:
		fmt.Printf("%T:", h, h)
	}
}

func main() {

	sa := secretAgent{
		person:  person{"Ivan", "Ivanovich"},
		country: "Russia",
	}

	fmt.Println(sa)
	fmt.Printf("sa type = %T\n", sa)
	sa.speak()
	fmt.Println("//------  --------------")

	bar(sa)

	p1 := person{"Gigi", "Robinet"}
	bar(p1)
}
