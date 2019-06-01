package main

import "fmt"

func main() {

	println("// 1 ------------------------------")
	type person struct {
		first   string
		last    string
		flavors []string
	}

	p1 := person{first: "Ana", last: "Banana", flavors: []string{"mere", "banane"}}
	p2 := person{first: "Gigi", last: "robinet", flavors: []string{"no idea", "12 23"}}
	fmt.Println(p1)
	fmt.Println(p2)
	for _, f := range p1.flavors {
		fmt.Println("\t", f)
	}

	println("// 2 ------------------------------")
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, p := range m {
		fmt.Println("person w/ key:", k)
		for _, f := range p.flavors {
			fmt.Println("\t+", f)
		}
	}

	fmt.Println(m)

	println("// 3 ------------------------------")

	type Vehicle struct {
		doors int
		color string
	}

	type Truck struct {
		Vehicle
		fourWheels bool
	}

	type Sedan struct {
		Vehicle
		luxury bool
	}

	vt := Truck{Vehicle: Vehicle{doors: 2, color: "pink"}, fourWheels: true}
	fmt.Println(vt, vt.doors)
	vs := Sedan{Vehicle: Vehicle{doors: 4, color: "cyan"}, luxury: true}
	fmt.Println(vs, vs.doors)

	println("// 4 ------------------------------")

	// anonymous structs

	type project struct {
		name     string
		location []float32
	}
	z := struct {
		code          string
		name          string
		public        bool
		sliceProjects []project
		mapProjects   map[string]project
	}{
		code:   "BLI",
		name:   "Barcoding Long Island",
		public: false,
		sliceProjects: []project{
			project{name: "Best thing", location: []float32{40.49, -73.04}},
			project{name: "OKish project", location: []float32{40.40, -73.23}},
		},
		mapProjects: map[string]project{
			"AAA": project{name: "Best thing", location: []float32{40.49, -73.04}},
			"BBB": project{name: "OKish project", location: []float32{40.40, -73.23}},
		},
	}
	fmt.Println(z)
	if !z.public {
		fmt.Println("+", z.code)
		fmt.Println("+", z.name)
		fmt.Println("+ slice projects:")
		for _, p := range z.sliceProjects {
			fmt.Println("\t", p)
		}
		fmt.Println("+ map projects:")
		for code, p := range z.mapProjects {
			fmt.Println("\t+", code, p)
		}
	}
}
