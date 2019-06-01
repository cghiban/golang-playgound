package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func buildHashKey(args []string) string {
	input := ""
	idx := 0
	// 1st pass, find the input
	for i := 1; i < len(args); i++ {
		if args[i] == "-query" {
			idx = i
			break
		}
	}

	if idx > 0 && idx < (len(args)-1) && args[idx+1] != "" {
		fmt.Printf("input file=%s\n", args[idx+1])
		input = args[idx+1]
	} else {
		fmt.Print("Error: missing input file")
		os.Exit(1)
	}

	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Println("Error: Input file not found")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("found input [%s] at index [%d]\n", input, idx)

	// 2nd pass, compute hash key
	h := md5.New()
	for i := 1; i < len(args); i++ {
		if i != idx && i != (idx+1) {
			io.WriteString(h, args[i])
		}
	}
	// now add the file
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	key := string(h.Sum(nil))
	return key
}
func main() {
	key := buildHashKey(os.Args)
	fmt.Printf("\nmd5sum: %x\n", key)
}
