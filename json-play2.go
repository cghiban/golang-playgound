package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Cornel",
		Last:  "Cornel",
		Age:   33,
	}
	p2 := person{
		First: "Mickey",
		Last:  "Mouse",
		Age:   99,
	}

	people := []person{p1, p2}

	fmt.Printf("%+v\n", people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bs))

}
