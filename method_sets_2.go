package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area:", s.area())
}

func main() {
	c := circle{5}

	fmt.Printf("%+v\n", c)
	//info(c) // this does not work
	info(&c)                     // this works
	fmt.Printf("%f\n", c.area()) // this works
}
