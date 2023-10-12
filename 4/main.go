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
	scores := []int{1, 2, 3, 5, 6}
	fmt.Println(scores)
	fmt.Println(scores[0:])  // first index - last index
	fmt.Println(scores[2:3]) // second index - third index
	fmt.Println(scores[:2])  // first index - second index
}
