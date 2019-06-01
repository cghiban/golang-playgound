package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// stringifier function
func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person
type ByName []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	xi := []int{9, 0, 14, 22, -1}
	xs := []string{"banana", "mere", "are", "Q", "B"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("-------------------------------")
	fmt.Println(xs, "sorted:", sort.StringsAreSorted(xs))
	sort.Strings(xs)
	fmt.Println(xs, "sorted:", sort.StringsAreSorted(xs))
	fmt.Println("-------------------------------")

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	//people = append(people, Person{"Michael", 11})
	sort.Sort(ByName(people))
	fmt.Println(people)
}
