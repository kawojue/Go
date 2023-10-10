package main

import "fmt"

func main() {
	var emptyArr []int

	anyArr := []any{20, "Code!"}

	var ages [3]int = [3]int{16, 17, 19}

	anyArr = append(anyArr, "Eat", "Sleep")

	fmt.Println(ages)
	fmt.Println(emptyArr)
	fmt.Printf("Array: %v\nLength: %d\n", anyArr, len(anyArr))

	// slices
	scores := []int{}
	fmt.Println(scores)
}
