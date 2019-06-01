package main

import "fmt"

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	x := append(s, "d")
	fmt.Println("xxx:", s, x)
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	z := []string{"A", "B", "C"}
	t := []string{"g", "h", "i"}
	t = append(t, "abc")
	t = append(t, z...) // append everuthing from z to t
	fmt.Println("dcl:", t)
	t = append(t[0:3], t[4:]...) // deleting "abc" from the slice t
	fmt.Println("after delete:", t)
	/*for i, v := range t {
		fmt.Printf("t[%d] = %v\n", i, v)
	}*/

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//------------------------
	for i := 'A'; i <= 'D'; i++ {
		fmt.Printf("%#U\n", i)
	}
	//------------------------
	//var buffer [30]byte
	var buffer = make([]byte, 10, 12) // (type, length, capacity)
	fmt.Println("len(buffer) = ", len(buffer))
	fmt.Println("cap(buffer) = ", cap(buffer))
	slice := buffer[0:10]
	//slice[10] = 99 // index out of range
	// but we can append to it
	slice = append(slice, 90)
	slice = append(slice, 91)
	fmt.Println("len(slice) = ", len(slice))
	fmt.Println("cap(slice) = ", cap(slice))
	slice = append(slice, 92)
	fmt.Println("len(slice) = ", len(slice))
	fmt.Println("cap(slice) = ", cap(slice))
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)

	//-------------------
}
