package main

import "fmt"

func one() {
	//var a [5]int
	x := [7]int{1, 3, 7, 11, 13, 17, 19}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func two() {
	fmt.Println("//-------------------------------")
	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, v := range y {
		fmt.Printf("%d ", v)
	}
	fmt.Println("")
	fmt.Printf("%T\n", y)

	fmt.Println(y[0:4])
	fmt.Println(y[1:5])
	fmt.Println(y[2:6])
	fmt.Println(y[3:7])
	//----
	fmt.Println("")
	for i := 0; i < 4; i++ {
		fmt.Println(y[i : i+4])
	}
}

// append to a slice
func four() {
	fmt.Println("//-------------------------------")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 55, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

// delete from a slice
func five() {
	fmt.Println("//-------------------------------")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[0:3], x[6:]...)
	fmt.Println(y)

}

func six() {
	fmt.Println("//-------------------------------")
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`,
		` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println(states)
	fmt.Printf("len=%d\tcap=%d\n", len(states), cap(states))
}

func six2() {
	fmt.Println("//-------------------------------")
	states := make([]string, 50, 100)
	fmt.Println(states)
	fmt.Printf("len=%d\tcap=%d\n", len(states), cap(states))

	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`,
		` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	for i := 0; i < len(states); i++ {
		fmt.Printf("%d\t %s\n", i, states[i])
	}
	//fmt.Println(states)
	fmt.Printf("len=%d\tcap=%d\n", len(states), cap(states))
	states = append(states, "xx", "yy")
	fmt.Printf("len=%d\tcap=%d\n", len(states), cap(states))
}

func seven() {
	fmt.Println("//- 7 ------------------------------")
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	twodslice := [][]string{x1, x2}

	fmt.Println(twodslice)
	for i, r := range twodslice {
		fmt.Println(i, r)
		for _, p := range twodslice[i] {
			fmt.Println("\t", p)
		}
	}
}

func eight() {
	fmt.Println("//- 8 ------------------------------")
	supermap := map[string][]string{
		"gigi":     []string{"a", "b", "c", "d"},
		"babaruba": []string{"fasole", "mămăligă", "chiperi"},
		"no_dr":    []string{"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(supermap)

	for k, v := range supermap {
		fmt.Println(k, v)
		for i, x := range v {
			fmt.Println("\t", i, x)
		}
	}

	fmt.Println("//- 9 ------------------------------")
	supermap["fănel"] = []string{"too many to list here"}
	supermap["babanu"] = []string{""}

	for k, v := range supermap {
		fmt.Println(k, v)
		/*for i, x := range v {
			fmt.Println("\t", i, x)
		}*/
	}

	fmt.Println("//- 9 ------------------------------")
	delete(supermap, "gigi")
	fmt.Println(supermap)
}

func main() {

	one()
	two()
	four()
	five()
	six()
	six2()
	seven()
	eight()
}
