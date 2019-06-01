package main

import (
	"bytes"
	"fmt"
)

func main() {
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			fmt.Println(i, j)
			if i < j {
				break
			}
		}
	}

	s := "something"
	ss := s[0:4]
	var buffer = []byte(s)
	bs := buffer[0:4]
	fmt.Println(s, buffer)
	fmt.Println(ss, bs)
	ePos := string.IndexRune(s, 'e')
	fmt.Println(ePos)
	ePos = bytes.IndexRune(bs, 'e')
	fmt.Println(ePos)
}
