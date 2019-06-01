package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Cornel","Last":"Ghiban","Age":33},{"First":"V","Last":"W","Age":123}]`

	bs := []byte(s)
	fmt.Printf("%T\n", s)
	//fmt.Println(s)
	fmt.Printf("%T\n", bs)
	//fmt.Println(bs)

	//people := []person{}
	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all data: ", people)
	for _, v := range people {
		fmt.Println(v)
	}
}
