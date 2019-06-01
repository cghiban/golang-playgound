package main

import "fmt"

func main() {

	m := map[string]int{
		"James":          33,
		"Miss Moneypeny": 28,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	// ok is a bool
	v, ok := m["Gigel"]
	fmt.Println(v, ok)

	if v, ok := m["Gigel"]; ok {
		fmt.Println("Found Gigel:", v)
	} else {
		fmt.Println("Not Found Gigel", v)
	}

	// iter the map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// add element
	m["Vasea"] = 99
	fmt.Println("after added: ", m)

	// delete
	delete(m, "Vasea")
	fmt.Println("after deleted: ", m)
}
