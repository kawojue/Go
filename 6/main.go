package main

import (
	"fmt"
	"slices"
	"sort"
)

type Person struct {
	Name string
	Age  uint8
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	scores := []int{5, 6, 1, 3}
	slices.Sort(scores)

	fmt.Println(scores)

	people := []Person{
		{"Raheem", 20},
		{"Precios", 21},
		{"My baby", 21},
	}

	sort.Sort(ByAge(people))

}
