package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log.SetOutput(os.Stderr)
	log.Println("okish")
	var msg string = "ok"
	var a string = `Gigi says: "ok"`
	fmt.Println(msg)
	fmt.Println(a)

	var i, j int
	i++
	fmt.Print(i, j)
	fmt.Printf(" // type of i is %T\n", i)
	fmt.Printf(" // %v -- %#v\n", i, i)
	ss := fmt.Sprintf("[[ %v -- %#v]]\n", i, i)
	fmt.Println(ss)

	type dna int
	var aa dna = 33
	fmt.Print(aa)
	fmt.Printf(" // type of dna is %T\n", aa)
	fmt.Printf(" // type of dna is %T\n", aa+1)
	//fmt.Printf(" // type of dna is %T\n", aa+i)

	j = int(aa)
	fmt.Print(j)
	fmt.Printf(" // type of j is %T\n", j)
	var s string
	i++
	fmt.Printf("\a[%s] is of type %T\n", s, s)

	//i = "hehe"  // this is not working here
	//fmt.Println(i, j)
	//fmt.Sprintf()
}
