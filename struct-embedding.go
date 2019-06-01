package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	jb := person{
		first: "James",
		last:  "Bond",
		age:   66,
	}
	agent := secretAgent{
		person: jb,
		ltk:    true,
	}

	miaua := person{
		first: "Miau",
		last:  "Mieww",
	}

	fmt.Printf("%v - %T\n", agent, agent)
	fmt.Printf("%v - %T\n", miaua, miaua)
	//fmt.Println(agent, miaua)

	fmt.Println(agent.first, agent.last)
	fmt.Println(miaua.first, miaua.last)
}
