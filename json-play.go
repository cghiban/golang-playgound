package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Person struct {
		//ID            int
		Name, address string
		Cars          []string
	}
	p := Person{
		Name:    "A B",
		address: "home",
		Cars:    []string{"Audi", "BMW", "Ford"},
	}

	fmt.Printf("%+v\n", p)

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println(string(b))
	os.Stdout.Write(b)
	fmt.Println("")
	fmt.Println("---------------------")

	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)
}
