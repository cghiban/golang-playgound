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

func main() {

	sa := secretAgent{
		person:  person{"Ivan", "Ivanovich"},
		country: "Russia",
	}

	fmt.Println(sa)
	sa.person.speak()
	sa.speak()
	fmt.Println("//------  --------------")
}
